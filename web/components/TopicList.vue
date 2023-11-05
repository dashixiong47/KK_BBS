<template>
  <div class="pb-5">
    <ul class="m-1">
      <li v-for="item in list" class="shadow-center w-full p-5 rounded-2xl mb-5">
        <Topic :detail="item" />
      </li>
    </ul>
    <Pagination v-model="pages" />
  </div>
</template>

<script setup>
import { useGetTopicList } from "@/api/server";
let props = defineProps({
  userId: {
    type: String,
    default: "",
  },
});
console.log(props,"-------------");
let pages = ref({
  total: 0,
  currentPage: 1,
  pageSize: 10,
});
const list = ref([]);
async function init() {
  try {
    const { data } = await useGetTopicList({ ...pages, userId: props.userId });

    list.value = data.list;
    pages.value.total = data.total;
  } catch (error) {}
}

init();
</script>
