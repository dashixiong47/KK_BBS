<template>
  <ul class="h-full m-1">
    <li
      v-for="item in 10"
      class="glass w-full p-5 rounded-2xl mb-5 shadow-center"
    >
      <Topic />
    </li>
  </ul>
</template>

<script setup>
const list = [
  {
    id: 1,
    title: "标题",
    content: "内容",
    type: "topic",
  },
];
</script>
