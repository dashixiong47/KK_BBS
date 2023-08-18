<template>
  <client-only>
    <div ref="editorContainer"></div>
  </client-only>
</template>

<script setup>
import { ref, onMounted, getCurrentInstance } from "vue";
let EditorJS,ImageTool,TelegramPost;
if (process.client) {
  EditorJS = (await import("@editorjs/editorjs")).default;
  ImageTool = (await import("@editorjs/image")).default;
  TelegramPost = (await import("editorjs-telegram-post")).default;
}
let { proxy } = getCurrentInstance();
let editor = ref(null);
onMounted(() => {
  setTimeout(() => {
    editor.value = new EditorJS({
      holder: proxy.$refs.editorContainer,
      i18n: {
        direction: "zh-cn",
      },
      tools: {
        image: {
          class: ImageTool,
          config: {
            endpoints: {
              byFile: "http://localhost:8008/uploadFile", // Your backend file uploader endpoint
              byUrl: "http://localhost:8008/fetchUrl", // Your endpoint that provides uploading by Url
            },
          },
        },
        telegramPost: TelegramPost,
      },
    });
  }, 1000);
});
</script>
