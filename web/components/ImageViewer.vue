<template>
  <Teleport to="body">
    <div
      class="image-viewer fixed z-10 top-0 left-0 inset-0 flex items-center justify-center bg-[--bg-masking-color] dark:bg-[--dark-bg-masking-color]"
      v-if="isOpen"
    >
      <!-- 覆盖层，点击关闭图片查看器 -->
      <!-- <div class="overlay absolute inset-0" @click="close"></div> -->

      <!-- 图片模态框 -->
      <div
        class="modal relative cursor-grab active:cursor-grabbing"
        ref="imageElement"
        @mousedown="handleMouseDown"
        @mousemove="handleMouseMove"
        @mouseup="handleMouseUp"
        @mouseleave="handleMouseUp"
      >
        <!-- 图片展示区域，应用缩放样式 -->
        <img :src="getPath(images[currentIndex].url)" :style="imageStyle" />
      </div>
      <!-- 控制按钮 -->
      <div
        class="controls flex absolute bottom-10 bg-[#606266] opacity-80 py-2 px-5 rounded-3xl"
      >
        <button @click="zoomIn">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="32"
            height="32"
            viewBox="0 0 1024 1024"
          >
            <path
              fill="#ffffff"
              d="m795.904 750.72l124.992 124.928a32 32 0 0 1-45.248 45.248L750.656 795.904a416 416 0 1 1 45.248-45.248zM480 832a352 352 0 1 0 0-704a352 352 0 0 0 0 704zm-32-384v-96a32 32 0 0 1 64 0v96h96a32 32 0 0 1 0 64h-96v96a32 32 0 0 1-64 0v-96h-96a32 32 0 0 1 0-64h96z"
            />
          </svg>
        </button>
        <button @click="zoomOut">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="32"
            height="32"
            viewBox="0 0 1024 1024"
          >
            <path
              fill="#ffffff"
              d="m795.904 750.72l124.992 124.928a32 32 0 0 1-45.248 45.248L750.656 795.904a416 416 0 1 1 45.248-45.248zM480 832a352 352 0 1 0 0-704a352 352 0 0 0 0 704zM352 448h256a32 32 0 0 1 0 64H352a32 32 0 0 1 0-64z"
            />
          </svg>
        </button>
        <button @click="resetTransform">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="32"
            height="32"
            viewBox="0 0 1024 1024"
          >
            <path
              fill="#ffffff"
              d="M771.776 794.88A384 384 0 0 1 128 512h64a320 320 0 0 0 555.712 216.448H654.72a32 32 0 1 1 0-64h149.056a32 32 0 0 1 32 32v148.928a32 32 0 1 1-64 0v-50.56zM276.288 295.616h92.992a32 32 0 0 1 0 64H220.16a32 32 0 0 1-32-32V178.56a32 32 0 0 1 64 0v50.56A384 384 0 0 1 896.128 512h-64a320 320 0 0 0-555.776-216.384z"
            />
          </svg>
        </button>
        <button @click="rotateLeft">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="32"
            height="32"
            viewBox="0 0 1024 1024"
          >
            <path
              fill="#ffffff"
              d="M289.088 296.704h92.992a32 32 0 0 1 0 64H232.96a32 32 0 0 1-32-32V179.712a32 32 0 0 1 64 0v50.56a384 384 0 0 1 643.84 282.88a384 384 0 0 1-383.936 384a384 384 0 0 1-384-384h64a320 320 0 1 0 640 0a320 320 0 0 0-555.712-216.448z"
            />
          </svg>
        </button>
        <button @click="rotateRight">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="32"
            height="32"
            viewBox="0 0 1024 1024"
          >
            <path
              fill="#ffffff"
              d="M784.512 230.272v-50.56a32 32 0 1 1 64 0v149.056a32 32 0 0 1-32 32H667.52a32 32 0 1 1 0-64h92.992A320 320 0 1 0 524.8 833.152a320 320 0 0 0 320-320h64a384 384 0 0 1-384 384a384 384 0 0 1-384-384a384 384 0 0 1 643.712-282.88z"
            />
          </svg>
        </button>
        <button @click="close">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="32"
            height="32"
            viewBox="0 0 1024 1024"
          >
            <path
              fill="#ffffff"
              d="M764.288 214.592L512 466.88L259.712 214.592a31.936 31.936 0 0 0-45.12 45.12L466.752 512L214.528 764.224a31.936 31.936 0 1 0 45.12 45.184L512 557.184l252.288 252.288a31.936 31.936 0 0 0 45.12-45.12L557.12 512.064l252.288-252.352a31.936 31.936 0 1 0-45.12-45.184z"
            />
          </svg>
        </button>
      </div>
      <button
        @click="previousImage"
        class="bg-[#606266] flex items-center justify-center absolute top-1/2 left-5 transform -translate-y-1/2 z-10 rounded-full h-14 w-14"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="32"
          height="32"
          viewBox="0 0 24 24"
        >
          <path
            fill="#ffffff"
            d="m14 18l-6-6l6-6l1.4 1.4l-4.6 4.6l4.6 4.6L14 18Z"
          />
        </svg>
      </button>
      <button
        @click="nextImage"
        class="bg-[#606266] flex items-center justify-center absolute top-1/2 right-5 transform -translate-y-1/2 z-10 rounded-full h-14 w-14"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="32"
          height="32"
          viewBox="0 0 24 24"
        >
          <path
            fill="#ffffff"
            d="M12.6 12L8 7.4L9.4 6l6 6l-6 6L8 16.6l4.6-4.6Z"
          />
        </svg>
      </button>
    </div>
  </Teleport>
</template>

<script setup>
const { getPath } = usePath();
// 图片数组、当前索引、缩放级别、图片元素引用
const images = ref([]);
const currentIndex = ref(0);
const scale = ref(1);
const imageElement = ref(null);
const isOpen = ref(false);
const translate = ref({ x: 0, y: 0 });
let isDragging = false;
let startX = 0;
let startY = 0;
const rotate = ref(0); // 图片的旋转角度

const handleMouseDown = (event) => {
  event.preventDefault(); // 阻止默认行为
  isDragging = true;
  startX = event.clientX - translate.value.x;
  startY = event.clientY - translate.value.y;
};

const handleMouseMove = (event) => {
  if (!isDragging) return;
  translate.value.x = event.clientX - startX;
  translate.value.y = event.clientY - startY;
};

const handleMouseUp = () => {
  isDragging = false;
};
// 重置图片到默认大小和位置
const resetTransform = () => {
  scale.value = 0.8;
  translate.value = { x: 0, y: 0 };
  rotate.value = 0;
};

// 左旋转图片
const rotateLeft = () => {
  rotate.value -= 90;
};

// 右旋转图片
const rotateRight = () => {
  rotate.value += 90;
};
// 打开图片查看器
const open = (newImages, newIndex) => {
  if (!newImages || !newImages.length) return;
  console.log(newImages);
  images.value = newImages;
  currentIndex.value = newIndex || 0;
  scale.value = 0.8;
  isOpen.value = true;
  setTimeout(() => {
    init();
  }, 0);
};

// 关闭图片查看器
const close = () => {
  isOpen.value = false;
};

// 放大图片
const zoomIn = () => {
  scale.value *= 1.1;
};

// 缩小图片
const zoomOut = () => {
  scale.value /= 1.1;
};

// 计算图片的样式
const imageStyle = computed(() => ({
  transform: `scale(${scale.value}) translate(${translate.value.x}px, ${translate.value.y}px) rotate(${rotate.value}deg)`,
  transition: isDragging ? "none" : "transform 0.3s ease",
}));
// 触摸事件的开始和结束位置
let touchStartX = 0;
let touchEndX = 0;

// 处理触摸开始事件
const handleTouchStart = (event) => {
  touchStartX = event.touches[0].clientX;
};

// 处理触摸移动事件
const handleTouchMove = (event) => {
  touchEndX = event.touches[0].clientX;
};

// 处理触摸结束事件
const handleTouchEnd = () => {
  if (touchEndX < touchStartX - 40) {
    nextImage();
  } else if (touchEndX > touchStartX + 40) {
    previousImage();
  }
};

// 显示下一张图片
const nextImage = () => {
  if (currentIndex.value < images.value.length - 1) {
    currentIndex.value++;
  } else {
    currentIndex.value = 0;
  }
};

// 显示上一张图片
const previousImage = () => {
  if (currentIndex.value > 0) {
    currentIndex.value--;
  } else {
    currentIndex.value = images.value.length - 1;
  }
};

// 处理鼠标滚轮事件
const handleWheel = (event) => {
  console.log(event);
  if (event.deltaY < 0) {
    zoomIn();
  } else {
    zoomOut();
  }
};

// 组件挂载时添加事件监听
function init() {
  if (process.client && imageElement.value) {
    const el = imageElement.value;
    el.addEventListener("wheel", handleWheel);
    el.addEventListener("touchstart", handleTouchStart);
    el.addEventListener("touchmove", handleTouchMove);
    el.addEventListener("touchend", handleTouchEnd);
  }
}

// 组件卸载时移除事件监听
onUnmounted(() => {
  if (process.client && imageElement.value) {
    const el = imageElement.value;
    el.removeEventListener("wheel", handleWheel);
    el.removeEventListener("touchstart", handleTouchStart);
    el.removeEventListener("touchmove", handleTouchMove);
    el.removeEventListener("touchend", handleTouchEnd);
  }
});

defineExpose({ open, close });
</script>
<style scoped>
.controls > button {
  margin: 0 10px;
}
.controls > button > svg {
  width: 20px;
  height: 20px;
}
</style>
