export const lazyload = {
    beforeMount(el, binding) {
      const src = binding.value;
  
      // 设置 div 的初始样式或类
      el.classList.add('placeholder');
  
      const img = new Image();
  
      img.onload = () => {
        // 当图片加载完毕后，替换 div 占位符
        el.replaceWith(img);
      };
  
      const observer = new IntersectionObserver(entries => {
        entries.forEach(entry => {
          if (entry.isIntersecting) {
            img.src = src;
            observer.unobserve(el);
          }
        });
      });
  
      observer.observe(el);
    }
  };
  