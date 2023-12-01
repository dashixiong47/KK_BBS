<template>
  <el-upload
    :show-file-list="false"
    class="border w-44 h-44 rounded-xl flex items-center justify-center"
    :before-upload="beforeUpload"
    v-loading="loading"
  >
    <img v-if="imageUrl" :src="getPath(imageUrl)" class="avatar" />
    <div v-else class="w-44 h-44 flex items-center justify-center">
      <el-icon><Plus /></el-icon>
    </div>
  </el-upload>
</template>

<script setup>
import { ref, computed } from "vue";
import { upload } from "@/api/upload.js";
import { getPath } from "@/utils/path.js";
const props = defineProps({
  modelValue: {
    type: String,
    default: "",
  },
});
let imageUrl = computed({
  get() {
    return props.modelValue;
  },
  set(val) {
    emit("update:modelValue", val);
  },
});
let loading = ref(false);
let emit = defineEmits(["update:modelValue"]);
const beforeUpload = async (file) => {
  const isJPG = file.type === "image/jpeg";
  const isLt2M = file.size / 1024 / 1024 < 2;
  loading.value = true;
  let formData = new FormData();
  formData.append("file", file);
  try {
    let { data } = await upload(formData);
    imageUrl.value = data.url;
    emit("update:modelValue", data.url);
  } catch (error) {
    console.log(error);
  }
  loading.value = false;
  return false;
};
</script>

<style lang="scss" scoped></style>
