"use strict";

const RenderObjectProto = {
    render(gl, programInfo, modelViewMatrix) {
        this.renderQueue.forEach(renderStep => {
            modelViewMatrix = renderStep(this, gl, programInfo, modelViewMatrix);
        });

        this.finishRender(this, gl, programInfo, modelViewMatrix);
    },

    update(dtSeconds) {
        this.updateQueue.forEach(updateStep => updateStep(this, dtSeconds));
    }
};

/**
 * Creates a render object composed of the various render components
 * @param {Object[]} components An array of RenderComponents
 * @returns {RenderObject}
 * @constructor
 */
const RenderObject = (components) => {
    let instance = Object.assign(Object.create(RenderObjectProto), {
        renderComponents: [],
        renderQueue: [],
        updateQueue: []
    });

    components.forEach(component => component.compose(instance));

    if (!instance.finishRender || typeof instance.finishRender !== "function") {
        throw new TypeError("No finishRenderFn was provided by any renderComponent of this renderObject");
    }

    return instance;
};

module.exports = RenderObject;
