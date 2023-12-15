<template>
  <div
    ref="imageContainer"
    class="w-full overflow-hidden"
    :style="{ height: props.height }"
  ></div>
</template>

<script setup>
import { onMounted, ref, watch } from "vue";

const props = defineProps({
  source: {
    type: String,
    required: true,
  },
  rolling: {
    type: Boolean,
    default: false,
  },
});
const imageContainer = ref(null);
const emit = defineEmits(["show"]);

const loadImage = () => {
  const img = new Image();
  img.onload = () => {
    img.className = "object-cover block";
    imageContainer.value.appendChild(img);
    imageContainer.value.style.height = "auto"; // 移除最小高度限制
  };
  img.src = props.source;
};

const setupIntersectionObserver = () => {
  const observer = new IntersectionObserver(
    (entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          if (!imageContainer.value.firstChild) {
            loadImage();
          }
          if (props.rolling) {
            emit("show");
          } else {
            observer.unobserve(imageContainer.value);
          }
        }
      });
    },
    { threshold: [0.7] }
  );
  observer.observe(imageContainer.value);
};

onMounted(() => {
  setupIntersectionObserver();
});

watch(
  () => props.source,
  (newSource) => {
    if (imageContainer.value.firstChild) {
      imageContainer.value.removeChild(imageContainer.value.firstChild);
    }
    loadImage();
  }
);
</script>

<style>
/* 你可以在这里添加额外的样式 */
.placeholder {
  /* 初始占位样式 */
}
</style>
