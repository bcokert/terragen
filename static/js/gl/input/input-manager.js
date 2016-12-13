"use strict";

var EventEmitter3 = require("eventemitter3");

const InputManagerProto = {
    KEY_TO_CODE: {
        "backspace": 8,
        "tab": 9,
        "enter": 13,
        "shift": 16,
        "ctrl": 17,
        "alt": 18,
        "escape": 27,
        "space": 32,
        "left": 37,
        "up": 38,
        "right": 39,
        "down": 40,
        "f1": 112,
        "f2": 113,
        "f3": 114,
        "f4": 115,
        "f5": 116,
        "f6": 117,
        "f7": 118,
        "f8": 119,
        "f9": 120,
        "f10": 121,
        "f11": 122,
        "f12": 123
    },

    setState(key, value) {
        if (typeof key === "number") {
            throw new Error("Tried to setState with a number in InputManager: " + key);
        }
        this.inputState[key] = value;
    },

    getState() {
        var newState = {};
        var keys = Object.keys(this.inputState);
        for (var i = 0; i < keys.length; i++) {
            newState[keys[i]] = this.inputState[keys[i]];
        }

        return newState;
    },

    emit(name, args) {
        this.emitter.emit(name, this.getState(), args);
    },

    on(name, callback) {
        this.emitter.on(name, callback);
    }
};

var letterKeys = " ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-_+=!@#$%^&*()<,>.?/|\\{[}]~`:;\"'";
for (var i = 0; i < letterKeys.length; i++) {
    InputManagerProto.KEY_TO_CODE[letterKeys[i]] = letterKeys.charCodeAt(i);
}

/**
 * An InputManager is essentially an observer with extra features, primarily related to aggregating calls and keeping track of the current input state.
 * Any client with access to the observer can register to listen for events, or tell it to fire events.
 * They can also subscribe to "builtin" events, like keydown, mousedown, etc
 * The purpose of builtin events is to remove the need for the client to add events for subtypes of other events
 * For example, instead of listening for keypress', determining if they're the right key, storing the state, firing a custom event,
 * and then listening for that event, builtins allow the client to simply listen for keypresses in general, and look at the state of the input manager on press
 * @returns {InputManager}
 * @constructor
 */
const InputManager = () => {
    const instance = Object.assign(Object.create(InputManagerProto), {
        emitter: new EventEmitter3(),
        inputState: {}
    });

    window.addEventListener("keydown", e => {
        if (!instance.inputState[e.keyCode]) {
            instance.inputState[e.keyCode] = Date.now();
            instance.emit("keydown", e.keyCode);
        }
    });

    window.addEventListener("keyup", e => {
        instance.inputState[e.keyCode] = null;
        instance.emit("keyup", e.keyCode);
    });

    window.addEventListener("mousedown", e => {
        instance.KEY_TO_CODE["mousedown"] = Date.now();
        instance.emit("mousedown", {x: e.clientX, values: e.clientY});
    });

    window.addEventListener("mousemove", e => {
        instance.KEY_TO_CODE["mousemove"] = Date.now();
        instance.emit("mousemove", {x: e.clientX, values: e.clientY});
    });

    window.addEventListener("mouseup", e => {
        instance.KEY_TO_CODE["mouseup"] = Date.now();
        instance.emit("mouseup", {x: e.clientX, values: e.clientY});
    });

    return instance;
};

InputManager.prototype = InputManagerProto;
InputManager.KEY_TO_CODE = InputManagerProto.KEY_TO_CODE;

module.exports = InputManager;
