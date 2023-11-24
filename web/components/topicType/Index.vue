<template>
  <KTabs :tabs="types" @selectd="selectd"></KTabs>
</template>

<script setup>
const types = [
  {
    name: "帖子",
    type: "topic",
    index: 1,
  },

  {
    name: "图片",
    type: "image",
    index: 2,
  },
  {
    name: "视频",
    type: "video",
    index: 3,
  },
  {
    name: "文章",
    type: "text",
    index: 4,
  },
];
let emit = defineEmits(["selectd"]);
const selectd = (index) => {
  emit("selectd", types[index]);
};
</script>
