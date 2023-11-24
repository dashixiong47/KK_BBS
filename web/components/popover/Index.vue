<template>
  <tippy :visible="getTooltipVisible">
    <span><slot></slot></span>
    <template #content>
      <slot name="content"></slot>
    </template>
  </tippy>
</template>

<script setup>
import { ref } from "vue";
import { Tippy,setDefaultProps } from "vue-tippy";

const porps = defineProps({
  modelValue: { type: Boolean, default: false },
});

const getTooltipVisible = computed({
  get: () => {
    console.log(porps.modelValue);
    return porps.modelValue;
  },
  set: (bl) => {
    emits("update:modelValue", bl);
  },
});

// 设置 Tippy 的默认属性
setDefaultProps({
  interactive: true,
  trigger: "manual", // 使用 manual 触发器，因为我们想手动控制显示和隐藏
  appendTo: document.body,
  theme: "glass",
});
</script>
