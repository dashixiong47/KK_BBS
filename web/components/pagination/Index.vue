<template>
  <div class="flex flex-col">
    <KButton
      class="w-full mb-5"
      :class="{
        disableColor: modelValue.page === total,
      }"
      :disabled="modelValue.page === total"
      @click="change(modelValue.page + 1)"
    >
      下一页
    </KButton>
    <div class="w-full flex justify-end">
      <div class="flex">
        <!-- 上一页按钮 -->
        <button
          class="shadow-center uncheckedColor mr-2 w-10 h-10 flex items-center justify-center rounded-lg cursor-pointer"
          :class="{
            disableColor: modelValue.page === 1,
            'cursor-not-allowed': modelValue.page === 1,
          }"
          :disabled="modelValue.page === 1"
          @click="change(modelValue.page - 1)"
        >
          <Icon name="iconoir:nav-arrow-left" />
        </button>

        <!-- 显示省略号 -->
        <button
          v-if="shouldShowEllipsis && modelValue.page > 4"
          class="shadow-center uncheckedColor mr-2 w-10 h-10 flex items-center justify-center rounded-lg cursor-pointer"
          @click="moveLeft"
        >
          ...
        </button>
        <!-- 显示中间的页数按钮 -->
        <button
          v-for="item in pages"
          :key="item"
          class="shadow-center uncheckedColor mr-2 w-10 h-10 flex items-center justify-center rounded-lg cursor-pointer"
          :class="{
            activeColor: item === modelValue.page,
          }"
          @click="change(item)"
        >
          {{ item }}
        </button>

        <!-- 显示省略号 -->
        <button
          v-if="shouldShowEllipsis && total - modelValue.page >= 4"
          class="shadow-center uncheckedColor mr-2 w-10 h-10 flex items-center justify-center rounded-lg cursor-pointer"
          @click="moveRight"
        >
          ...
        </button>

        <!-- 下一页按钮 -->
        <button
          class="shadow-center uncheckedColor w-10 h-10 flex items-center justify-center rounded-lg cursor-pointer"
          :class="{
            disableColor: modelValue.page === total,
            'cursor-not-allowed': modelValue.page === total,
          }"
          :disabled="modelValue.page === total"
          @click="change(modelValue.page + 1)"
        >
          <Icon name="iconoir:nav-arrow-right" />
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, reactive } from "vue";

const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({
      total: 0,
      pageSize: 10,
      page: 1,
    }),
  },
});
const modelValue = computed(() => props.modelValue);
const total = computed(() => {
  return modelValue.value.total
    ? Math.ceil(modelValue.value.total / modelValue.value.pageSize)
    : 1;
});
const shouldShowEllipsis = computed(() => total.value > 9); // 根据总页数决定是否显示省略号
const emit = defineEmits(["update:modelValue", "query"]);

function change(index) {
  modelValue.value.page = index;
  emit("update:modelValue", modelValue.value);
  emit("query");
}

function moveLeft() {
  change(Math.max(1, modelValue.value.page - 5)); // 向左移动5页，但不超过第1页
}

function moveRight() {
  change(Math.min(total.value, modelValue.value.page + 5)); // 向右移动5页，但不超过最后一页
}

const pages = computed(() => {
  let start, end;
  if (total.value <= 9) {
    start = 1;
    end = total.value;
  } else {
    const sideButtons = 3; // 当总页数大于9时，每边显示3个按钮
    start = modelValue.value.page - sideButtons;
    end = modelValue.value.page + sideButtons;

    if (start < 1) {
      end += 1 - start; // 如果起始页码小于1，则相应增加结束页码
      start = 1; // 确保起始页码不小于1
    }
    if (end > total.value) {
      start -= end - total.value; // 如果结束页码大于总页数，则相应减少起始页码
      end = total.value; // 确保结束页码不超过总页数
    }
  }
  return Array.from({ length: end - start + 1 }, (_, index) => start + index); // 生成页码数组
});
</script>
