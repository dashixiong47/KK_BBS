<template>
  <div ref="colorPickerContainer"></div>
</template>

<script setup>
let colorPickerContainer = ref(null);
let emit = defineEmits(["change"]);
let props = defineProps({
  color: {
    type: String,
    default: () => {
      return "rgba(255, 255, 255, 1)";
    },
  },
});
let color = computed(() => props.color);
let colorPicker = null;
onMounted(() => {
  showColorPicker();
});
watch(
  () => color.value,
  (val) => {
    colorPicker.color.set(val);
  }
);
function showColorPicker() {
  colorPicker = new iro.ColorPicker(colorPickerContainer.value, {
    color: color.value,
    width: 100,
    layout: [
      {
        component: iro.ui.Slider,
        options: {
          sliderType: "red", // 红色滑块
        },
      },
      {
        component: iro.ui.Slider,
        options: {
          sliderType: "green", // 绿色滑块
        },
      },
      {
        component: iro.ui.Slider,
        options: {
          sliderType: "blue", // 蓝色滑块
        },
      },
      {
        component: iro.ui.Slider,
        options: {
          sliderType: "alpha", // 透明度滑块
        },
      },
      {
        component: iro.ui.Wheel, // 仅使用颜色轮组件
      },
    ],
    layoutDirection: "horizontal", // 垂直布局
  });
  colorPicker.on("color:change", function (color) {
    let rgbaColor = color.rgba;
    const rgbaString = `rgba(${rgbaColor.r}, ${rgbaColor.g}, ${rgbaColor.b}, ${rgbaColor.a})`;
    emit("change", rgbaString);
  });
}
</script>
