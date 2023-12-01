<template>
  <div class="pb-5">
    <component
      class="mb-5"
      :is="typeOptions[props.type]"
      :list="list"
    ></component>

    <Pagination v-model="pages" @query="init" />
  </div>
</template>

<script setup>
import { useGetTopicList } from "@/api/server";
import { useGroupStore } from "~/stores/init.js";
import Img from "./Img.vue";
import Text from "./Text.vue";
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
    type: Number,
    default: "",
  },
});
let pages = ref({
  total: 0,
  page: 1,
  pageSize: 10,
});
let typeOptions = {
  2: Img,
  4: Text,
};
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
