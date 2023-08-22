<template>
  <div class="relative">
    <button
      class="relative z-10 bg-blue-500 hover:bg-blue-700 text-white font-bold w-10 h-10 rounded-full"
      @click="change"
    >
      <Icon name="icon-zhutise" />
    </button>
    <transition
      name="scale-transition"
      enter-active-class="animate-scaleUpFromRightBottom"
      leave-active-class="animate-scaleDownToRightBottom"
    >
      <div
        v-if="showPopup"
        class="glass p-5 rounded-2xl w-80 h-96 !absolute right-0 bottom-0 origin-bottom-right"
      >
        <slot />
      </div>
    </transition>
  </div>
</template>

<script setup>
import { watch, ref } from "vue";
import { inject } from "vue";

let showPopup = ref(false);
let actived = inject("actived");

let { index } = defineProps({
  index: {
    type: Number,
    default: 0,
  },
});
let emit = defineEmits(["change"]);
function change() {
  showPopup.value = !showPopup.value;
  actived.value = index;
  emit("change", showPopup.value);
}
watch(
  () => actived.value,
  (val) => {
    showPopup.value = val === index;
  }
);
</script>
