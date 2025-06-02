import axios from "axios";

// initializing axios
const api = axios.create({
  baseURL: "localhost:8080",
});

function determineChunkSize(fileSize) {
  const minChunkSize = 20 * 1024 * 1024; // 20 MB
  const maxChunkSize = 100 * 1024 * 1024; // 100 MB
  const maxParts = 410;

  let chunkSize = Math.ceil(fileSize / maxParts);

  // Clamp the chunk size between min and max
  chunkSize = Math.max(chunkSize, minChunkSize);
  chunkSize = Math.min(chunkSize, maxChunkSize);

  return chunkSize;
}

// original source: https://github.com/pilovm/multithreaded-uploader/blob/master/frontend/uploader.js
export class Uploader {
  constructor(options) {
    this.useTransferAcceleration = options.useTransferAcceleration;
    // this must be bigger than or equal to 5MB,
    // otherwise AWS will respond with:
    // "Your proposed upload is smaller than the minimum allowed size"
    options.chunkSize = options.chunkSize || 0;
    this.chunkSize = determineChunkSize(options.file.size)
  
    // number of parallel uploads
    options.threadsQuantity = options.threadsQuantity || 0;
    this.threadsQuantity = 15;
    // adjust the timeout value to activate exponential backoff retry strategy
    this.timeout = 0;
    this.file = options.file;

    this.fileName = options.fileName;
    this.aborted = false;
    this.uploadedSize = 0;
    this.progressCache = {};
    this.activeConnections = {};
    this.parts = [];
    this.uploadedParts = [];
    this.fileId = null;
    this.fileKey = null;
    this.onProgressFn = () => {};
    this.onErrorFn = () => {};
    this.baseURL = "http://localhost:8080";
  }

  start() {
    this.initialize();
  }

  async initialize() {
    try {
      // adding the the file extension (if present) to fileName
      let fileName = this.file.name;

      // initializing the multipart request
      const videoInitializationUploadInput = {
        key: "video.mp4",
        bucket_name: "testbucketkab",
      };
      const initializeReponse = await api.request({
        url: "/api/v1/createupload",
        method: "POST",
        data: videoInitializationUploadInput,
        baseURL: this.baseURL,
      });

      const AWSFileDataOutput = initializeReponse.data;
      console.log(AWSFileDataOutput);
      this.fileId = AWSFileDataOutput.uploadId;
      //   this.fileKey = AWSFileDataOutput.fileKey

      // retrieving the pre-signed URLs
      const numberOfparts = Math.ceil(this.file.size / this.chunkSize);
    console.log(numberOfparts)
      const parts = [];

      for (let i = 1; i <= numberOfparts; i++) {
        parts.push(i);
      }
      const AWSMultipartFileDataInput = {
        key: "video.mp4",
        bucket: "testbucketkab",
        upload_id: this.fileId,
        part_number: parts,
      };

      const urlsResponse = await api.request({
        url: "/api/v1/getPresign",
        method: "POST",
        data: AWSMultipartFileDataInput,
        baseURL: this.baseURL,
      });

      const newParts = urlsResponse.data.data;

      console.log(newParts);
      this.parts.push(...newParts);
      this.parts.sort((a, b) => a.part_number - b.part_number);
      this.parts.shift();
      this.sendNext();
      console.log(this.parts);
    } catch (error) {
      await this.complete(error);
    }
  }

  sendNext(retry = 0) {
    console.log("started");
    const activeConnections = Object.keys(this.activeConnections).length;

    if (activeConnections >= this.threadsQuantity) {
      console.log("startedds");
      return;
    }

    if (!this.parts.length) {
      if (!activeConnections) {
        this.complete();
      }

      return;
    }

    const part = this.parts.pop();
    if (this.file && part) {
      const sentSize = (part.part_number - 1) * this.chunkSize;
      const chunk = this.file.slice(sentSize, sentSize + this.chunkSize);
      console.log("starteewed");
      const sendChunkStarted = () => {
        this.sendNext();
      };

      this.sendChunk(chunk, part, sendChunkStarted)
        .then(() => {
          this.sendNext();
        })
        .catch((error) => {
          if (retry <= 6) {
            retry++;
            const wait = (ms) => new Promise((res) => setTimeout(res, ms));
            //exponential backoff retry before giving up
            console.log(
              `Part#${part.part_number} failed to upload, backing off ${
                2 ** retry * 100
              } before retrying...`
            );
            wait(2 ** retry * 100).then(() => {
              this.parts.push(part);
              this.sendNext(retry);
            });
          } else {
            console.log(`Part#${part.part_number} failed to upload, giving up`);
            this.complete(error);
          }
        });
    }
  }

  async complete(error) {
    if (error && !this.aborted) {
      this.onErrorFn(error);
      return;
    }

    if (error) {
      this.onErrorFn(error);
      return;
    }

    try {
      await this.sendCompleteRequest();
    } catch (error) {
      this.onErrorFn(error);
    }
  }

  async sendCompleteRequest() {
    if (this.fileId) {
      const videoFinalizationMultiPartInput = {
        upload_id: this.fileId,
        key: "video.mp4",
        bucket: "testbucketkab",
        parts: this.uploadedParts.sort((a, b) => a.part_number - b.part_number),
      };

      await api.request({
        url: "/api/v1/completeupload",
        method: "POST",
        data: videoFinalizationMultiPartInput,
        baseURL: this.baseURL,
      });
    }
  }

  sendChunk(chunk, part, sendChunkStarted) {
    return new Promise((resolve, reject) => {
      this.upload(chunk, part, sendChunkStarted)
        .then((status) => {
          if (status !== 200) {
            reject(new Error("Failed chunk upload"));
            return;
          }

          resolve();
        })
        .catch((error) => {
          reject(error);
        });
    });
  }

  handleProgress(part, event) {
    if (this.file) {
      if (
        event.type === "progress" ||
        event.type === "error" ||
        event.type === "abort"
      ) {
        this.progressCache[part] = event.loaded;
      }

      if (event.type === "uploaded") {
        this.uploadedSize += this.progressCache[part] || 0;
        delete this.progressCache[part];
      }

      const inProgress = Object.keys(this.progressCache)
        .map(Number)
        .reduce((memo, id) => (memo += this.progressCache[id]), 0);

      const sent = Math.min(this.uploadedSize + inProgress, this.file.size);

      const total = this.file.size;

      const percentage = Math.round((sent / total) * 100);

      this.onProgressFn({
        sent: sent,
        total: total,
        percentage: percentage,
      });
    }
  }

  upload(file, part, sendChunkStarted) {
    // uploading each part with its pre-signed URL
    return new Promise((resolve, reject) => {
      console.log("iploading", part);
      const throwXHRError = (error, part, abortFx) => {
        delete this.activeConnections[part.part_number - 1];
        reject(error);
        window.removeEventListener("offline", abortFx);
      };
      if (this.fileId) {
        if (!window.navigator.onLine) reject(new Error("System is offline"));

        const xhr = (this.activeConnections[part.part_number - 1] =
          new XMLHttpRequest());
        xhr.timeout = this.timeout;
        sendChunkStarted();

        const progressListener = this.handleProgress.bind(
          this,
          part.part_number - 1
        );

        console.log("iploadingw", part);

        xhr.upload.addEventListener("progress", progressListener);

        xhr.addEventListener("error", progressListener);
        xhr.addEventListener("abort", progressListener);
        xhr.addEventListener("loadend", progressListener);

        xhr.open("PUT", part.url);
        const abortXHR = () => xhr.abort();
        xhr.onreadystatechange = () => {
          if (xhr.readyState === 4 && xhr.status === 200) {
            const ETag = xhr.getResponseHeader("ETag");

            if (ETag) {
              const uploadedPart = {
                part_number: part.part_number,
                etag: ETag.replaceAll('"', ""),
              };

              this.uploadedParts.push(uploadedPart);

              resolve(xhr.status);
              delete this.activeConnections[part.part_number - 1];
              window.removeEventListener("offline", abortXHR);
            }
          }
        };

        xhr.onerror = (error) => {
          throwXHRError(error, part, abortXHR);
        };
        xhr.ontimeout = (error) => {
          throwXHRError(error, part, abortXHR);
        };
        xhr.onabort = () => {
          throwXHRError(
            new Error("Upload canceled by user or system"),
            part,
            abortXHR
          );
        };
        window.addEventListener("offline", abortXHR);
        xhr.send(file);
      }
    });
  }

  onProgress(onProgress) {
    this.onProgressFn = onProgress;
    return this;
  }

  onError(onError) {
    this.onErrorFn = onError;
    return this;
  }

  abort() {
    Object.keys(this.activeConnections)
      .map(Number)
      .forEach((id) => {
        this.activeConnections[id].abort();
      });

    this.aborted = true;
  }
}
