<template>
  <div class="relative">
    <div class="relative z-10" @click="change">
      <slot />
    </div>
    <transition
      name="scale-transition"
      enter-active-class="animate-scaleUpFromRightBottom"
      leave-active-class="animate-scaleDownToRightBottom"
    >
      <div
        v-if="showPopup"
        class="glass p-5 rounded-2xl w-80 h-96 !absolute right-0 bottom-0 origin-bottom-right"
      >
        <slot name="content" />
      </div>
    </transition>
  </div>
</template>

<script setup>
let showPopup = ref(false);
function change() {
  console.log("------");
  showPopup.value = !showPopup.value;
}
</script>
