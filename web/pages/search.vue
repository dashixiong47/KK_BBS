<template>
  <TopicType @selectd="selectdTopicType" />
  <ul class="my-5 flex">
    <li v-for="item in getGroup">
      <KButton
        class="mr-5 flex items-center"
        @click="selectdGroup(item)"
        :actived="queryData.group === item.id"
      >
        {{ item.name }}
      </KButton>
    </li>
  </ul>
  <ul>
    <li v-for="item in dataList" class="shadow-center rounded-xl p-4 mb-5">
      <KLink :to="getUrl(item)">
        <div class="flex max-h-40">
          <div class="w-[80%] overflow-hidden">
            <h3 class="flex items-center">
              标题：
              <div class="text-sm" v-html="item.title"></div>
            </h3>
            <div
              v-html="item.content"
              class="mt-2 truncate text-xs text-[--font-secondary-color] dark:text-[--dark-font-secondary-color]"
            ></div>
          </div>

          <div
            class="w-[18%] ml-[2%] flex items-center justify-center rounded-xl overflow-hidden border border-cyan-500"
          >
            <KImage
              v-if="item.cover"
              class="w-full h-full"
              :source="getPath(JSON.parse(item.cover)[0])"
            ></KImage>
            <span v-else>无封面</span>
          </div>
        </div>
      </KLink>
    </li>
  </ul>
</template>

<script setup>
import { useGroupStore } from "~/stores/init.js";

import { useSearch } from "@/api/server";
const { getPath } = usePath();
let store = useGroupStore();
let getGroup = computed(() => [
  {
    id: -1,
    name: "全部",
  },
  ...store.getGroup,
]);
const route = useRoute();

let queryData = ref({
  q: route.query.q,
  type: -1,
  group: -1,
});
let type = {
  1: "topic",
  2: "img",
  3: "video",
  4: "text",
};
let dataList = ref([]);
async function init() {
  try {
    let { data } = await useSearch(queryData.value);
    dataList.value = data.list;
    console.log(data);
  } catch (error) {}
}
function selectdTopicType(val) {
  queryData.value = {
    ...queryData.value,
    type: val.index,
  };
  init();
}
function selectdGroup(val) {
  queryData.value = {
    ...queryData.value,
    group: val.id,
  };
  init();
}
function getUrl(item) {
  return `/${type[item.type]}/${item.id}`;
}
watch(
  () => route.query.q,
  (newV) => {
    queryData.value = {
      ...queryData.value,
      q: newV,
    };
    init();
  }
);
store.fetchGroup();
init();
</script>

<style lang="postcss">
em {
  color: red;
}
</style>
