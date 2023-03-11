function animate(obj, target, callback) {
    // console.log(callback); callback = function() {} 调用的时候 callback()
    // 先清除以前的定时器，只保留当前的一个定时器执行
    clearInterval(obj.timer);
    obj.timer = setInterval(function() {
        // 步长值写到定时器的里面，现在的位置obj.offsetLeft
        // 我们把步长值改为整数，不要让他出现小数问题,
        // var step = Math.ceil((target - obj.offsetLeft) / 10);
        var step = (target - obj.offsetLeft) / 10;
        // 注意回退时的问题，往上取整(Math.ceil)，往下取整(Math.floor)
        step = step > 0 ? Math.ceil(step) : Math.floor(step);
        if (obj.offsetLeft == target) {
            // 停止动画 本质是停止定时器
            clearInterval(obj.timer);
            // 回调函数写在定时器结束里面的后面
            if (callback) {
                // 调用函数
                callback();
            }
        }
        // 把每次加1 这个步长值改为一个慢慢变小的值  步长公式：(目标值 - 现在的位置) / 10
        obj.style.left = obj.offsetLeft + step + 'px';

    }, 15);
}
