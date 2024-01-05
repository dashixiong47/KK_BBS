<template>
  <div ref="imageContainer" class="w-full h-full overflow-hidden">
    <!-- 显示默认图片，直到真实图片加载完成 -->
    <img
      v-if="showDefault"
      src="/images/loading.png"
      class="w-full h-full object-cover block"
    />
  </div>
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
  fillContainer: {
    type: Boolean,
    default: true,
  },
});
const imageContainer = ref(null);
const emit = defineEmits(["show"]);
let show = ref(false);
let showDefault = ref(true); // 状态控制默认图片的显示

const loadImage = () => {
  const img = new Image();
  img.onload = () => {
    img.className = props.fillContainer
      ? "w-full h-full object-cover block"
      : "max-w-full max-h-full";
    imageContainer.value.appendChild(img);
    showDefault.value = false; // 图片加载后隐藏默认图片
  };
  img.src = props.source;
};

const setupIntersectionObserver = () => {
  const observer = new IntersectionObserver(
    (entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          if (show.value === false) {
            loadImage();
          }
          show.value = true;
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
    showDefault.value = true; // 重置默认图片的显示状态
    loadImage();
  }
);
</script>
