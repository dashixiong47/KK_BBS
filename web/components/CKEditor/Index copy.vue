<template>
  <client-only>
    <ckeditor
      :editor="ClassicEditor"
      v-model="editorData"
      :config="editorConfig"
    ></ckeditor>
  </client-only>
</template>

<script setup>
let CKEditor;
let ClassicEditor;
if (process.client) {
  CKEditor = (await import("@ckeditor/ckeditor5-vue")).default;
  ClassicEditor = (await import("@ckeditor/ckeditor5-build-classic")).default;
  await import("@ckeditor/ckeditor5-build-classic/build/translations/zh-cn.js");
}
const editorData = ref("");
const editorConfig = ref({
  language: "zh-cn",
  extraPlugins: [MyCustomUploadAdapterPlugin],
});
class MyUploadAdapter {
  constructor(loader) {
    // The file loader instance to use during the upload.
    this.loader = loader;
  }

  // Starts the upload process.
  upload() {
    return new Promise((resolve, reject) => {
      // Here you should use the file loader to send the file to your server.
      // On success, call resolve() with an object that has a "default" property with the URL of the uploaded file.
      // On error, call reject() with an error message.

      // Example:
      const data = new FormData();
      data.append("upload", this.loader.file);
      resolve({
            default:"https://images.hxsj.in/test/1.jpg",
          });
    //   fetch("", {
    //     method: "POST",
    //     body: data,
    //   })
    //     .then((response) => response.json())
    //     .then((data) => {
    //       resolve({
    //         default: data.url,
    //       });
    //     })
    //     .catch(reject);
    });
  }

  // Aborts the upload process.
  abort() {
    // If you want to support aborting, set the abort controller and call abort() on it.
  }
}
function MyCustomUploadAdapterPlugin(editor) {
  editor.plugins.get("FileRepository").createUploadAdapter = (loader) => {
    // Configure and return the adapter.
    return new MyUploadAdapter(loader);
  };
}
</script>
<style scoped>
:deep(.ck-editor__editable) {
  height: 800px !important;
}
</style>
