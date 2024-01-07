<template>
  <div class="pb-5">
    <div
      v-if="!list"
      class="w-full h-full text-[--font-secondary-color] flex flex-col items-center justify-center p-10"
    >
      <Icon name="ep:box" size="10rem"></Icon>
      <span class="mt-2">无数据</span>
    </div>
    <div v-if="loading">
      <Skeleton class="mb-5" :lines="5" v-for="item in 10"></Skeleton>
    </div>
    <component
      v-else
      class="mb-5"
      :is="typeOptions[props.type]"
      :list="list"
    ></component>
    <Pagination v-model="pages" @query="init" />
  </div>
</template>

<script setup>
import { useGetTopicList, useCollect } from "~/api/server";
import { useGroupStore } from "~/stores/init.js";
import Topic from "./Topic.vue";
import Image from "./Image.vue";
import Text from "./Text.vue";
import Video from "./Video.vue";
let loading = ref(false);
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
    default: -1,
  },
  initFunc: {
    type: Number,
    default: 0,
  },
});
let pages = ref({
  total: 0,
  page: 1,
  pageSize: 10,
});
let typeOptions = {
  "-1": Topic,
  1: Topic,
  2: Image,
  3: Video,
  4: Text,
};
watch(
  () => props.userId,
  (val) => {
    pages.value = {
      total: 0,
      page: 1,
      pageSize: 10,
    };
    init(val);
  }
);
watch(
  () => props.type,
  (val) => {
    console.log(val);
    pages.value = {
      total: 0,
      page: 1,
      pageSize: 10,
    };
    init(val);
  }
);
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
  loading.value = true;
  let func;
  switch (props.initFunc) {
    case 1:
      func = useCollect;
      break;
    default:
      func = useGetTopicList;
      break;
  }
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
    const { data = { list: [], total: 0 } } = await func(query);
    list.value = data.list;
    pages.value.total = data.total;
  } catch (error) {
    console.log(error);
  }
  loading.value = false;
}
</script>
