<template>
  <ul class="h-full m-1">
    <li v-for="item in list" class="glass w-full p-5 rounded-2xl mb-5">
      <Topic :detail="item" />
    </li>
  </ul>
</template>

<script setup>
import { useGetTopicList } from "@/api/server";
const list = ref([]);
async function init() {
  try {
    const { data } = await useGetTopicList({
      page: 1,
      pageSize: 10,
    });
    list.value = data;
  } catch (error) {}
}
init();
</script>
