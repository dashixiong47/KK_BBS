<template>
  <div @click="triggerFileUpload">
    <slot></slot>
    <input
      type="file"
      :accept="accept"
      :multiple="multiple"
      ref="fileInput"
      @change="handleFileUpload"
      style="display: none"
    />
  </div>
</template>

<script setup>
import { ref } from "vue";
import { upload } from "~/api/upload";
let { accept, multiple } = defineProps({
  accept: {
    type: String,
    default: "image/*",
  },
  multiple: {
    type: Boolean,
    default: false,
  },
});
const { addMessage } = useMessage();
// 创建一个 ref 用于访问文件输入元素
let fileInput = ref(null);
let emit = defineEmits(["uploadSuccess"]);

// 处理文件上传
const handleFileUpload = async () => {
  for (const file of fileInput.value.files) {
    try {
      await uploadFile(file);  // 等待文件上传成功后再继续
      // 可以在这里添加一些上传成功后的处理逻辑
    } catch (error) {
      // 处理上传过程中的错误
      console.error("上传失败: ", error);
      break; // 如果一个文件上传失败，跳出循环，停止上传后续文件
    }
  }
};


const uploadFile = async (file) => {
  try {
    let res = await upload(file, progress);
    emit("uploadSuccess", { ...res, type: file.type, size: file.size });
  } catch (error) {
    addMessage(error, "error");
  }
};
// 触发文件上传对话框
const triggerFileUpload = () => {
  fileInput.value.click();
};

const progress = (event) => {
  console.log(event);
};
</script>
