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

const warn = (isDebug, msg) => {
    if (isDebug) {
        console.warn(msg);
    }
}

const run = () => {
    if (!validateParams(window.ssiWe)) {
        console.error('ssi-we: Invalid params');
        return;
    }
    const params = mergeObjects({
        reconnectDelay: 1000,
        reconnectAttempts: 5,
        reconnect: true,
        openEvent: false,
        debug: false,
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
            const name = el.getAttribute('data-we-name');
            if (!name) {
                warn(params.debug, 'ssi-we: Element without name property. In the element: ' + el.tagName.toLowerCase());
                return;
            }
            el.addEventListener(event, (e) => {
                const p = window.ssiWe?.params || {};
                const paramsAttrs = el.getAttributeNames().filter((attr) => attr.startsWith('data-we-params-'));
                paramsAttrs.forEach((attr) => {
                    const key = attr.replace('data-we-params-', '');
                    p[key] = el.getAttribute(attr);
                });
                sendEvent(s, {
                    page: window.ssiWe.page,
                    type: e.type,
                    name: name,
                    element: el.tagName.toLowerCase(),
                    params: Object.keys(p).length > 0 ? p : undefined,
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
        s.send("Hello Server!")
        unListenAllEvents();
        listenAllEvents(s);
    }

    socket.onopen = () => {
        connection = true;
        if (params.openEvent) {
            sendEvent(socket, {
                page: window.ssiWe.page,
                type: 'open',
                name: 'open',
                element: 'window',
                params: window.ssiWe?.params,
                user: window.ssiWe?.user
            });
        }
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
ssiWeInit();