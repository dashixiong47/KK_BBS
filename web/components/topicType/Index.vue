<template>
   <KTabs :tabs="types" @selectd="selectd"></KTabs>
</template>

<script setup>
const types = [
  {
    name: "帖子",
    type: "topic",
  },
  {
    name: "视频",
    type: "video",
  },
  {
    name: "文章",
    type: "text",
  },
];
const selectd = (index) => {
  console.log(index);
};
</script>
