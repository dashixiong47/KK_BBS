<template>
  <div class="py-5">
    <!-- 使用已排序的响应式数据进行渲染 -->
    <div v-for="item in sortedTextContent" :key="item.order">
      <h2 class="text-xl font-bold mb-5">{{ item.name }}</h2>
      <pre v-html="item.data" class="whitespace-pre-wrap"></pre>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';
const { getPath } = usePath();
const props = defineProps({
  detail: {
    type: Object,
    default: () => {},
  },
});

let textContent = ref([]);

// 计算属性，用于排序textContent
const sortedTextContent = computed(() => {
  return textContent.value.sort((a, b) => a.order - b.order);
});

// 初始化函数
async function init() {
  try {
    let list = props.detail.texts.map((item) => getText(item));
    textContent.value = await Promise.all(list);
  } catch (error) {
    console.error("初始化数据时发生错误:", error);
    // 这里可以添加更多错误处理逻辑
  }
}

// 异步获取文本数据
async function getText(item) {
  try {
    let response = await useFetch("/api/text", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        url: getPath(item.url),
      }),
    });
    return {
      name: item.name,
      order: item.order,
      data: response.data.value,
    };
  } catch (error) {
    console.error("获取文本数据时发生错误:", error);
    throw error; // 抛出错误，以便在调用链中处理
  }
}

// 调用初始化函数
init();
</script>
