import React, { useState, useRef } from "react";
import { useDropzone } from "react-dropzone";
import { CloudUpload, XCircle } from "lucide-react";
import { Uploader } from "../utils";
const FileUploader = () => {
  const [file, setFile] = useState(null);
  const [progress, setProgress] = useState(0);
  const [uploading, setUploading] = useState(false);
    const [pgvalue, setPgvalue] = useState(undefined);
    const [perf, setPerf] = useState(undefined);
    const [baseUrl, setBaseUrl] = useState(undefined);
    const [partsize, setPartsize] = useState(undefined);
    const [numuploads, setNumuploads] = useState(undefined);
    const [ta, setTa] = useState(undefined);
  const controllerRef = useRef(null);

  const onDrop = (acceptedFiles) => {
    if (acceptedFiles.length === 0) return;
    const selected = acceptedFiles[0];
    setFile(selected);
    const uploaderOptions = {
      file: selected,
     
      chunkSize: partsize,
      threadsQuantity: numuploads,
     
    };

    let percentage = undefined;
    setPgvalue(0);
    setPerf("-");
    const uploader = new Uploader(uploaderOptions);
    const tBegin = performance.now();
    uploader
      .onProgress(({ percentage: newPercentage }) => {
        // to avoid the same percentage to be logged twice
        if (percentage === 100) {
          setPerf((performance.now() - tBegin) / 1000);
        }
        if (newPercentage !== percentage) {
          percentage = newPercentage;
          setPgvalue(percentage);
        }
      })
      .onError((error) => {
        setFile(undefined);
        console.error(error);
      });

    uploader.start();      
  };

  const { getRootProps, getInputProps, isDragActive } = useDropzone({
    onDrop,
    multiple: false,
  });

  const uploadFile = async (file) => {
    const controller = new AbortController();
    controllerRef.current = controller;

    setUploading(true);
    setProgress(0);

    const interval = setInterval(() => {
      setProgress((prev) => {
        if (prev >= 100) {
          clearInterval(interval);
          setUploading(false);
        }
        return prev + 4;
      });
    }, 150);

    await new Promise((res) => setTimeout(res, 4000)); // simulate 4s upload
  };

  const abortUpload = () => {
    if (controllerRef.current) {
      controllerRef.current.abort();
      setUploading(false);
      setProgress(0);
    }
  };

  return (
    <div className="h-screen w-screen flex items-center justify-center bg-gray-50 px-4">
      <div className="w-full max-w-md bg-white rounded-xl shadow-lg p-6">
        <div
          {...getRootProps()}
          className={`border-2 border-dashed rounded-lg p-8 flex flex-col items-center justify-center transition-colors duration-200 cursor-pointer ${
            isDragActive
              ? "border-blue-500 bg-blue-50"
              : "border-gray-300 bg-gray-100 hover:bg-gray-200"
          }`}
        >
          <input {...getInputProps()} />
          <CloudUpload className="h-10 w-10 text-gray-500 mb-2" />
          <p className="text-gray-700 text-sm">
            {isDragActive
              ? "Drop your file here"
              : "Drag & drop a file here, or click to browse"}
          </p>
        </div>

        {file && (
          <div className="mt-6">
            <p className="text-sm text-gray-800 font-medium truncate">
              {file.name}
            </p>

            <div className="w-full mt-2 bg-gray-200 rounded-full h-3 overflow-hidden">
              <div
                className="bg-blue-500 h-full transition-all duration-300"
                style={{ width: `${pgvalue}%` }}
              ></div>
              <div>
                {perf}
              </div>
            </div>

            {uploading ? (
              <button
                onClick={abortUpload}
                className="mt-3 flex items-center gap-1 text-red-600 text-sm hover:underline"
              >
                <XCircle className="w-4 h-4" />
                Cancel Upload
              </button>
            ) : pgvalue === 100 ? (
              <p className="text-green-600 text-sm mt-2">✅ Upload complete!</p>
            ) : null}
          </div>
        )}
      </div>
    </div>
  );
};

export default FileUploader;
