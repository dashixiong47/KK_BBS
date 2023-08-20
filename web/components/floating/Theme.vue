<template>
  <Floating :index="index">
    <div>
      <label>透明度</label>
      <input
        type="range"
        :value="blurValue"
        @input="updateValue"
        min="0"
        max="20"
      />
    </div>
    <div>
      <label>背景色</label>
      <input type="range" @input="updateValue" min="0" max="20" />
    </div>
  </Floating>
</template>

<script setup>
import { useThemeStore } from "~/stores/theme.js";
let themeStore = useThemeStore();
let config = computed(() => themeStore.getThemeConfig);
let blurValue = ref(0);
let { index } = defineProps({
  index: {
    type: Number,
    default: 0,
  },
});
useHead({
  style: [
    {
      type: "text/css",
      innerHTML: `:root { --blur-value: ${config.value.blur}px !important; }
      `,
    },
  ],
});

onMounted(() => {
  initTheme();
  const rootStyles = getComputedStyle(document.documentElement);
  const variableValue = rootStyles.getPropertyValue("--blur-value").trim();
  console.log(variableValue);
});
if (process.client) {
}
function initTheme() {
  blurValue.value = config.value.blur;
}
function updateValue() {
  const cssText = `--blur-value: ${event.target.value}px !important`;
  document.documentElement.style.cssText += cssText;
}
</script>
