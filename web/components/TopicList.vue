<template>
  <div class="pb-5">
    <ul class="m-1">
      <li v-for="item in list" class="glass w-full p-5 rounded-2xl mb-5">
        <Topic :detail="item" />
      </li>
    </ul>
    <Pagination v-model="pages" />
  </div>
</template>

<script setup>
import { useGetTopicList } from "@/api/server";
let pages = ref({
  total: 100,
  currentPage: 1,
  pageSize: 10,
});
const list = ref([]);
async function init() {
  try {
    const { data } = await useGetTopicList(pages);

    list.value = data.list;
    console.log(data);
    // pages.value.total = data.total;
  } catch (error) {}
}

init();
</script>
