<template>
  <div class="pt-5" ref="textBox">
    <!-- 使用已排序的响应式数据进行渲染 -->
    <div v-if="textContent[modelValue]">
      <h2 class="text-xl text-center font-bold mb-5">
        {{ textContent[modelValue].name }}
      </h2>
      <pre
        v-html="textContent[modelValue].data"
        class="whitespace-pre-wrap"
      ></pre>
      <div class="flex items-center justify-between mt-5">
        <KButton
          class="w-[45%]"
          :disabled="!modelValue"
          :class="{ disableColor: !modelValue }"
          @click="change(1)"
          >上一章</KButton
        >
        <KButton
          class="w-[45%]"
          :disabled="modelValue === detail.texts.length - 1"
          :class="{ disableColor: modelValue === detail.texts.length - 1 }"
          @click="change(2)"
          >下一章</KButton
        >
      </div>
    </div>
    <Skeleton v-else :lines="20" />
  </div>
</template>

<script setup>
const props = defineProps({
  detail: {
    type: Object,
    default: () => ({}),
  },
  modelValue: {
    type: Number,
    default: 0,
  }, // 当前选中的文章编号
});
const { getPath } = usePath();
const textContent = ref({});
const textBox = ref(null);
const emit = defineEmits(["update:modelValue"]);
// 异步获取文本数据
async function getText(item) {
  try {
    let res = await fetch(getPath(item.url), {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });
    let data = await res.text();
    return {
      name: item.name,
      order: item.order,
      data: data,
    };
  } catch (error) {
    console.error("获取文本数据时发生错误:", error);
    return {
      name: item.name,
      order: item.order,
      data: "获取失败，请重新获取",
    };
  }
}
function change(type) {
  let order = props.modelValue;
  if (type === 1) {
    order--;
  } else {
    order++;
  }

  emit("update:modelValue", order);
  show();
}
watch(
  () => props.modelValue,
  async (newVal) => {
    let info = props.detail.texts.find((item) => item.order === newVal);

    if (textContent.value[newVal]) {
      show();
      return;
    }
    textContent.value[newVal] = await getText(info);
  },
  {
    immediate: true,
  }
);
function show() {
  const element = textBox.value;
  element.scrollIntoView({ behavior: "smooth" });
}
</script>
