<template>
  <div class="w-full rounded-xl flex flex-wrap" ref="textList">
    <div
      :key="item.name + index"
      v-for="(item, index) in texts"
      class="viewer w-32 h-48 p-2 shadow-center relative rounded-xl mr-5 mb-7"
    >
      <div class="text-xs w-full h-full overflow-hidden">
        {{ textData[item.name] }}
      </div>
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
      class="w-32 h-48 flex items-center justify-center rounded-xl shadow-center active:shadow-active"
      @uploadSuccess="uploadSuccess"
      accept="text/plain"
      :multiple="true"
    >
      <div>+</div>
    </Upload>
  </div>
  <div class="flex justify-end">
    <KButton @click="toRepeat">去除重复</KButton>
    <KButton class="ml-5" @click="sort">一键排序</KButton>
  </div>
  <Dialog v-model="viewerVisible" class="w-[60%]">
    <pre class="w-96 max-h-96 overflow-y-auto whitespace-pre-wrap">
      {{ viewer }}
    </pre>
  </Dialog>
</template>

<script setup>
import Sortable from "sortablejs";
const { getPath } = usePath();
const { getSize } = useSize();
const { t } = useI18n();
let texts = ref([]);
let textData = ref({});
let textList = ref(null);
let viewerVisible = ref(false);
let viewer = ref("");
onMounted(() => {
  new Sortable(textList.value, {
    animation: 150,
    ghostClass: "sortable-ghost",
    onUpdate: async (evt) => {
      const oldIndex = evt.oldIndex;
      const newIndex = evt.newIndex;
      const movedItem = images.value.splice(oldIndex, 1)[0];
      images.value.splice(newIndex, 0, movedItem);
    },
  });
});
const uploadSuccess = async (res) => {
  let url = getPath(res.url);
  texts.value.push({
    name: res.name,
    url: url,
    size: getSize(res.size) + "M",
  });
  textData.value[res.name] = await getText(url);
};
const toRepeat = () => {
  let obj = {};
  texts.value = texts.value.reduce((cur, next) => {
    obj[next.url] ? "" : (obj[next.url] = true && cur.push(next));
    return cur;
  }, []);
};
const sort = () => {
  texts.value.sort((a, b) => {
    return a.name.localeCompare(b.name);
  });
};
const openViewer = (index) => {
  viewerVisible.value = true;
  viewer.value = textData.value[texts.value[index].name];
};
// 校验参数
const checkParams = () => {
  if (!texts.value.length) {
    return t("topic-create-text");
  }
  return false;
};
const remove = (index) => {
  texts.value.splice(index, 1);
};
// 异步获取文本数据
async function getText(url) {
  try {
    let response = await useFetch("/api/text", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        url,
      }),
    });
    return response.data.value;
  } catch (error) {
    console.error("获取文本数据时发生错误:", error);
    throw error; // 抛出错误，以便在调用链中处理
  }
}
// 提交
const getFormData = () => {
  let list = texts.value.map((item, index) => {
    return { ...item, order: index };
  });
  return {
    topicText: {
      texts: JSON.stringify(list),
    },
  };
};
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
