function ssiWeInit() {
    if (typeof window.ssiWe === 'undefined') {
        window.ssiWe = {
            url: 'localhost:8080',
            page: 'home',
            params: {},
            user: {}
        };
    }
    run();
}

const validateParams = (params) => {
    return !!params && typeof params === 'object' && params.url && params.page;
}

const mergeObjects = (obj1, obj2) => {
    return Object.assign({}, obj1, obj2);
}

const run = () => {
    if (validateParams(window.ssiWe)) {
        console.warn('ssi-we: Invalid params');
        return;
    }
    const params = mergeObjects({
        reconnectDelay: 1000,
        reconnectAttempts: 5,
        reconnect: true,
    }, window.ssiWe);
    let socket = new WebSocket(`ws://${window.ssiWe.url}/ws`);
    let connection = false;

    const sendEvent = (s, p) => {
        if (!connection) return;
        s.send(JSON.stringify(p));
    }

    const reconnect = () => {
        if (params.reconnect) {
            setTimeout(() => {
                socket = new WebSocket(`ws://${window.ssiWe.url}/ws`);
            }, params.reconnectDelay);
        }
    }

    const listenAllEvents = (s) => {
        const elements = document.querySelectorAll('[data-we-event]');
        elements.forEach((el) => {
            const event = el.getAttribute('data-we-event') || 'click';
            el.addEventListener(event, (e) => {
                const _p = el.getAttribute('data-we-params');
                const p = _p ? mergeObjects(window.ssiWe?.params, JSON.parse(_p)) : window.ssiWe?.params;
                sendEvent(s, {
                    page: window.ssiWe.page,
                    type: e.type,
                    element: el.tagName.toLowerCase(),
                    params: p,
                    user: window.ssiWe?.user
                });
            });
        });
    };

    const unListenAllEvents = () => {
        const elements = document.querySelectorAll('[data-we-event]');
        elements.forEach((el) => {
            const event = el.getAttribute('data-we-event') || 'click';
            el.removeEventListener(event, () => { });
        });
    }

    const onOpen = (s) => {
        unListenAllEvents();
        listenAllEvents(s);
    }

    socket.onopen = () => {
        connection = true;
        sendEvent(socket, {
            page: window.ssiWe.page,
            type: 'open',
            element: 'window',
            params: window.ssiWe?.params,
            user: window.ssiWe?.user
        });
        onOpen(socket);
    }

    socket.onclose = () => {
        connection = false;
        reconnect();
    }

    socket.onerror = () => {
        connection = false;
        reconnect();
    }
}