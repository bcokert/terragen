"use strict";

const TextureLoaderProto = {
    loadResources(gl, resources) {
        return new Promise((resolve) => {
            var loadedTextures = {};
            var finishLoading = (texture, textureImage, textureSource) => {
                gl.bindTexture(gl.TEXTURE_2D, texture);
                if (textureSource.invertY) {
                    gl.pixelStorei(gl.UNPACK_FLIP_Y_WEBGL, true);
                }
                gl.texImage2D(gl.TEXTURE_2D, 0, gl.RGBA, gl.RGBA, gl.UNSIGNED_BYTE, textureImage);
                gl.texParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST);
                gl.texParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST);
                gl.bindTexture(gl.TEXTURE_2D, null);

                loadedTextures[textureSource.name] = texture;

                if (Object.keys(loadedTextures).length === resources.length) {
                    resolve(loadedTextures);
                }
            };

            resources.forEach(textureSource => {
                var texture = gl.createTexture();
                var textureImage = new Image();
                textureImage.onload = () => finishLoading(texture, textureImage, textureSource);
                textureImage.src = textureSource.url;
            });
        });
    }
};

/**
 * A TextureLoader loads the given images, and creates gl textures from them
 * @returns {TextureLoader}
 * @constructor
 */
const TextureLoader = () => {
    return Object.assign(Object.create(TextureLoaderProto), {});
};

TextureLoader.prototype = TextureLoaderProto;

module.exports = TextureLoader;
