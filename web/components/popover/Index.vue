<template>
  <tippy>
    <span><slot></slot></span>
    <template #content class="bg-white">
      <slot name="content"></slot>
    </template>
  </tippy>
</template>

<script setup>
import { Tippy, setDefaultProps } from "vue-tippy";
setDefaultProps({
  placement: "top",
  interactive: true,
  trigger: "click",
  appendTo: document.body,
  theme: "glass",
});
function change() {
  console.log("change");
}
</script>
