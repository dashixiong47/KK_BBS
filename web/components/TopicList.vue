<template>
  <div class="pb-5">
    <ul class="m-1">
      <li
        v-for="item in list"
        class="shadow-center w-full p-5 rounded-2xl mb-5"
      >
        <Topic :detail="item" />
      </li>
    </ul>
    <Pagination v-model="pages" @query="init" />
  </div>
</template>

<script setup>
import { useGetTopicList } from "@/api/server";
import { useGroupStore } from "~/stores/init.js";
const store = useGroupStore();
const actived = computed(() => {
  return store.actived;
});
let props = defineProps({
  userId: {
    type: String,
    default: "",
  },
  type: {
    type: [Number, String],
    default: "",
  },
});
let pages = ref({
  total: 0,
  page: 1,
  pageSize: 10,
});
watch(
  () => actived.value,
  (val) => {
    pages.value = {
      total: 0,
      page: 1,
      pageSize: 10,
    };
    init(val);
  },
  {
    immediate: true,
  }
);
const list = ref([]);
async function init() {
  try {
    let query = {
      ...pages.value,

      type: props.type,
    };
    if (props.userId) {
      query["userId"] = props.userId;
    }
    if (actived.value) {
      query["groupId"] = actived.value;
    }
    const { data = { list: [], total: 0 } } = await useGetTopicList(query);
    list.value = data.list;
    pages.value.total = data.total;
  } catch (error) {
    console.log(error);
  }
}
</script>
