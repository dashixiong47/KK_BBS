<template>
  <div class="w-full rounded-xl flex flex-wrap" ref="imageList">
    <div
      :key="item.name + index"
      v-for="(item, index) in videos"
      class="viewer w-48 h-32 shadow-center relative rounded-xl mr-5 mb-7"
    >
      <!-- <PlayVideo :src="item.url"></PlayVideo> -->
      <img
        class="rounded-xl absolute object-cover w-full h-full"
        :src="item.url.replace('index.m3u8', 'cover.jpg')"
        alt=""
        srcset=""
      />
      <p class="text-sm text-center py-1 absolute -bottom-8 truncate w-full">
        {{ item.name }}
      </p>
      <div class="mask rounded-xl absolute top-0 right-0 w-full h-full">
        <Icon
          @click="openViewer(index)"
          class="cursor-pointer hover:text-[--illuminate-color]"
          size="1.5rem"
          name="ep:view"
        ></Icon>
        <Icon
          @click="remove(index)"
          class="ml-2 cursor-pointer hover:text-[--illuminate-color]"
          size="1.5rem"
          name="ep:delete"
        ></Icon>
      </div>
    </div>
    <Upload
      class="w-48 h-32 flex items-center justify-center rounded-xl shadow-center active:shadow-active"
      @uploadSuccess="uploadSuccess"
      :multiple="true"
      accept="video/*"
    >
      <div>+</div>
    </Upload>
  </div>
  <div class="flex justify-end">
    <KButton @click="toRepeat">去除重复</KButton>
    <KButton class="ml-5" @click="sort">一键排序</KButton>
  </div>
  <Dialog v-model="playVideoShow">
    <PlayVideo class="my-5 w-[800px]" :src="currentVideoUrl" />
  </Dialog>
</template>

<script setup>
import Sortable from "sortablejs";
const { getPath } = usePath();
const { t } = useI18n();
let videos = ref([]);
let imageList = ref(null);
onMounted(() => {
  new Sortable(imageList.value, {
    animation: 150,
    ghostClass: "sortable-ghost",
    onUpdate: async (evt) => {
      const oldIndex = evt.oldIndex;
      const newIndex = evt.newIndex;
      const movedItem = videos.value.splice(oldIndex, 1)[0];
      videos.value.splice(newIndex, 0, movedItem);
    },
  });
});
let playVideoShow = ref(false);
let currentVideoUrl = ref("");
const remove = (index) => {
  videos.value.splice(index, 1);
};
const uploadSuccess = (res) => {
  videos.value.push({
    name: res.name,
    url: getPath(res.url),
  });
};
const toRepeat = () => {
  let obj = {};
  videos.value = videos.value.reduce((cur, next) => {
    obj[next.url] ? "" : (obj[next.url] = true && cur.push(next));
    return cur;
  }, []);
};
const sort = () => {
  videos.value.sort((a, b) => {
    // return a.name.localeCompare(b.name, undefined, { numeric: true });
    return a.name.localeCompare(b.name);
  });
};
// 校验参数
const checkParams = () => {
  if (!videos.value.length) {
    return t("topic-create-video");
  }
  return false;
};

// 提交
const getFormData = () => {
  let list = videos.value.map((item, index) => {
    return { ...item, order: index };
  });
  return {
    topicVideo: {
      videos: JSON.stringify(list),
    },
  };
};
function openViewer(index) {
  playVideoShow.value = true;
  currentVideoUrl.value = videos.value[index].url;
}
defineExpose({
  getFormData,
  checkParams,
});
</script>

<style scoped>
.mask {
  display: none;
  background: rgba(0, 0, 0, 0.5);
  transition: all 0.3s;
  align-items: center;
  justify-content: center;
}
.viewer:hover .mask {
  display: flex;
}
</style>
