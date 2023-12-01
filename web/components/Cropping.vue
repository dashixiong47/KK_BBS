<template>
  <input
    class="hidden"
    type="file"
    ref="file"
    @change="onFileChange"
    :accept="acceptedTypesString"
  />
  <div @click="() => $refs.file.click()">
    <slot></slot>
  </div>
  <Dialog v-model="dialogVisible" :title="dialogTitle" @confirm="cropImage">
    <div ref="croppieRef" class="croppieRef"></div>
  </Dialog>
</template>

<script setup>
import { ref, watch, computed } from "vue";
import Croppie from "croppie";
import "croppie/croppie.css";
import { upload } from "~/api/upload";
const {addMessage} = useMessage()
const props = defineProps({
  cropWidth: { type: Number, default: 200 },
  cropHeight: { type: Number, default: 200 },
  cropType: { type: String, default: "circle" }, // 'square' or 'circle'
  dialogTitle: { type: String, default: "裁剪图片" },
  acceptedTypes: { type: Array, default: () => ["image/jpeg", "image/png"] },
});
const originalFileName = ref('');
const croppieRef = ref(null);
let croppieInstance = null;
const imageUrl = ref("");
let dialogVisible = ref(false);
let emit = defineEmits(["uploadSuccess"]);
// 将acceptedTypes数组转换为逗号分隔的字符串
const acceptedTypesString = computed(() => props.acceptedTypes.join(","));

watch(
  () => imageUrl.value,
  (newUrl) => {
    dialogVisible.value = !!newUrl;
    if (newUrl) {
      setTimeout(() => {
        croppieInstance = new Croppie(croppieRef.value, {
          url: newUrl,
          viewport: {
            width: props.cropWidth,
            height: props.cropHeight,
            type: props.cropType,
          },
          boundary: { width: 300, height: 300 },
        });
      }, 0);
    }
  }
);

const onFileChange = (e) => {
  const file = e.target.files[0];
  if (!props.acceptedTypes.includes(file.type)) {
    alert("不支持的文件类型！");
    return;
  }
  originalFileName.value = file.name; // 保存原始文件名
  const reader = new FileReader();
  reader.onload = (e) => {
    imageUrl.value = e.target.result;
  };
  reader.readAsDataURL(file);
};

const cropImage = () => {
  croppieInstance
    .result({
      type: "canvas",
      size: "viewport",
      format: "png", // 指定输出格式，根据需要调整
      quality: 1, // 设置质量，1为最高质量
    })
    .then((croppedImgBase64) => {
      // 将Base64转换为Blob
      const blob = base64ToBlob(croppedImgBase64, "image/jpeg");

      // 将Blob转换为File
      const file = new File([blob], originalFileName.value, {
        type: "image/jpeg",
      });

      // 在这里处理裁剪后的图片，如上传到服务器
      uploadFile(file);
    });
};

// Base64转Blob函数
function base64ToBlob(base64, mimeType) {
  const byteCharacters = atob(base64.split(",")[1]);
  const byteNumbers = new Array(byteCharacters.length);
  for (let i = 0; i < byteCharacters.length; i++) {
    byteNumbers[i] = byteCharacters.charCodeAt(i);
  }
  const byteArray = new Uint8Array(byteNumbers);
  return new Blob([byteArray], { type: mimeType });
}
const uploadFile = async (file) => {
  try {
    let res = await upload(file);
    emit("uploadSuccess", { ...res, type: file.type, size: file.size });
  } catch (error) {
    addMessage(error, "error");
  }
};
</script>
