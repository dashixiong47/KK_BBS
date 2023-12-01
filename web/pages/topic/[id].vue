<template>
  <div class="">
    <!-- 侧边状态栏 -->
    <Sidebar :detail="detail" />
    <div class="grid grid-cols-9 gap-5">
      <div class="relative h-full col-span-9 sm:col-span-6">
        <Card>
          <div class="flex" v-if="detail.user">
            <Avatar :url="detail.user?.avatar" />
            <div class="w-full ml-2 flex items-center justify-between">
              <div class="flex flex-col items-start">
                <KLink to="#" class="text-md font-bold font-main-color">
                  {{ detail.user?.nickname }}
                </KLink>
                <span class="text-xs font-regular-color">
                  发布时间:{{ getRelativeTime(detail.createdAt) }}
                </span>
              </div>
            </div>
          </div>
          <h2 class="font-medium text-xl mt-5">{{ detail.title }}</h2>
          <component :is="type[detail.type]" :detail="detail.topicDetail" />
        </Card>
        <Card id="comment" class="mt-5">
          <Comments :topicId="route.params.id" @change="commentChange" />
        </Card>
      </div>
      <div class="sm:col-span-3 sm:block m-1">
        <div class="sticky top-0">
          <Card>
            <p class="border-b pb-2 mb-2 font-bold">附件</p>
            <AttachmentList :topicId="route.params.id" />
          </Card>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useRoute } from "#app";
import { useGetTopicDetail } from "~/api/server";
import { useAppConfigStore } from "~/stores/init";
import Default from "~/components/detail/Index.vue";
import Images from "~/components/detail/Images.vue";
import Texts from "~/components/detail/Texts.vue";
let appConfigStore = useAppConfigStore();
let appConfig = computed(() => appConfigStore.getConfig);

const { getRelativeTime } = useTime();
const route = useRoute();

let detail = ref({});
let type = {
  1: Default,
  2: Images,
  4: Texts,
};
async function getTopicDetail() {
  try {
    // const {data} = await useFetch("http://localhost:8080/api/v1/topic/" + route.params.id);
    let { data } = await useGetTopicDetail(route.params.id);
    detail.value = data || {};
    useHead({
      title: detail.value.title + " - " + appConfig.value.appName,
    });
  } catch (error) {
    console.log(error);
    showError(error);
  }
}
function commentChange() {
  detail.value.comment = detail.value.comment + 1;
  detail.value.commentState = true;
}
function init() {
  getTopicDetail();
}
init();
</script>
