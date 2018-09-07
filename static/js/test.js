function b(a) {
    var c = new WebSocket(a);
    c.onclose = function () {
        setTimeout(function () {
            b(a)
        }, 2E3)
    };
    c.onmessage = function () {
        location.reload()
    }
}
