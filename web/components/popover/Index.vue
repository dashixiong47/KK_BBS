<template>
  <tippy>
    <span><slot></slot></span>
    <template #content>
      <slot name="content"></slot>
    </template>
  </tippy>
</template>

<script setup>
import { Tippy, setDefaultProps } from "vue-tippy";
let { placement } = defineProps({
  placement: {
    type: String,
    default: "top",
  },
});
setDefaultProps({
    placement: placement,
    interactive: true,
    trigger: "click",
    appendTo: document.body,
    theme: "glass",
  });
</script>
