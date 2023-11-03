<template>
  <div ref="skeleton" class="skeleton p-4">
    <div
      v-for="i in finalLines"
      :key="i"
      class="mb-2 bg-gray-300 animate-pulse rounded h-4 w-full"
    ></div>
  </div>
</template>

<script setup>

const props = defineProps({
  lines: {
    type: Number,
    default: 0,  // 默认值设置为0，表示未指定
  },
});
const { lines } = toRefs(props);

const skeleton = ref(null);  // 创建一个 ref 来引用 div 元素

let computedLines = ref(0);  // 创建一个 ref 来存储计算出的行数

onMounted(() => {
  if (lines.value === 0) {  // 只有在没有指定 lines 时才计算行数
      const parentHeight = skeleton.value.offsetHeight;  // 获取父元素的高度
      const lineHeight = 20;  // 假设每行的高度（包括 margin-bottom）是 20px
      computedLines.value = Math.floor(parentHeight / lineHeight);  // 计算所需的行数
  }
});

const finalLines = computed(() => {
  return lines.value > 0 ? lines.value : computedLines.value;  // 如果指定了 lines，就使用它，否则使用计算出的行数
});

</script>
