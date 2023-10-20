
export default {
  mounted(el, binding) {
    console.log(el);
    const handleScroll = () => {
      const scrollHeight = el.scrollHeight;
      const scrollTop = el.scrollTop;
      const clientHeight = el.clientHeight;

      if (scrollTop + clientHeight >= scrollHeight) {
        binding.value(); // 调用绑定的函数
      }
    };

    onMounted(() => {
      el.addEventListener('scroll', handleScroll);
    });

    onUnmounted(() => {
      el.removeEventListener('scroll', handleScroll);
    });
  }
};
