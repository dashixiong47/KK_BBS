<template>
  <div class="w-full h-full">
    <input
      type="file"
      ref="fileInput"
      @change="handleFileUpload"
      style="display: none"
    />
    <div class="w-full h-full" @click="triggerFileUpload">
      <slot></slot>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { upload } from "~/api/upload";
// 创建一个 ref 用于访问文件输入元素
let fileInput = ref(null);
let emit = defineEmits(["uploadSuccess"]);

// 处理文件上传
const handleFileUpload = async () => {
  const file = fileInput.value.files[0];
  try {
    let res = await upload(file, progress);
    let data = new URL(res);
    emit("uploadSuccess", {
      url: data.origin + data.pathname,
      id: data.searchParams.get("id"),
      name: data.searchParams.get("name"),
    });
  } catch (error) {}
};

// 触发文件上传对话框
const triggerFileUpload = () => {
  fileInput.value.click();
};

const progress = (event) => {
  console.log(event);
};
</script>
