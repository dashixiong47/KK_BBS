// composables/useNotice.js
import { ref, computed } from 'vue';

export default function useNotice() {
    let noticeCount = 0;
    let noticeContainer
    onMounted(() => {
        noticeContainer = document.createElement('div');
        noticeContainer.style.position = 'fixed';
        noticeContainer.style.top = '50px';
        noticeContainer.style.right = '5px';
        noticeContainer.style.zIndex = '50';
        document.body.appendChild(noticeContainer);
    })

    function notice(options) {
        const {
            title = '提示',
            content,
            autoClose = true
        } = options;

        const div = document.createElement('div');
        const html = `
        <i class="iconfont icon-close absolute top-4 right-4 text-xs cursor-pointer text-light-5"></i>
        <p class="text-md font-bold">${title}</p>
        <div class="text-sm">${content}</div>`;
        div.innerHTML = html;

        div.className = `glass text-dark-1 dark:text-light-1 transition-transform duration-300 ease-in-out p-4 rounded-xl overflow-hidden w-80 min-h-[80px] text-gray-900 shadow-md`;
        div.style.setProperty('--translate-y', `${noticeCount * 20}px`);
        div.classList.add('animate-slideInFromRight');
        div.dataset.position = noticeCount;

        if (autoClose) {
            let progressBar = document.createElement('div');
            progressBar.className = 'h-1 w-full bg-green-500 rounded-xl absolute bottom-0 left-0 animate-[widthToZero_5s_ease-in-out_forwards]';
            div.appendChild(progressBar);

            // 增加通知计数的倍数来调整超时时间
            setTimeout(() => {
                removeNotice(div);
            }, 5000 + (noticeCount * 500));  // 通知每隔0.5秒消失
        }

        noticeCount++;
        noticeContainer.appendChild(div);
        div.querySelector('.icon-close').onclick = () => {
            removeNotice(div);
        };
    }


    function removeNotice(div) {
        if (!noticeContainer.contains(div)) {
            return;
        }

        div.classList.remove('animate-slideInFromRight');
        div.classList.add('animate-slideOutToLeft');

        setTimeout(() => {
            const divPosition = parseInt(div.dataset.position, 10);
            noticeCount--;
            const index = Array.from(noticeContainer.children).indexOf(div);
            noticeContainer.removeChild(div);

            for (let i = index; i < noticeContainer.children.length; i++) {
                const el = noticeContainer.children[i];
                const currentPosition = parseInt(el.dataset.position, 10);
                if (currentPosition > divPosition) {
                    const currentTransformValue = parseInt(getComputedStyle(el).getPropertyValue('--translate-y'), 10);
                    el.style.setProperty('--translate-y', `${currentTransformValue - 20}px`);
                    el.dataset.position = currentPosition - 1;
                }
            }
        }, 300); // 这里的300ms是动画的持续时间，确保它与您的动画时间相匹配
    }

    return { notice };
}
