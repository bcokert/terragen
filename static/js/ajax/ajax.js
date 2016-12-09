"use strict";

class Ajax {
    /**
     * Returns a new url that has the given query params added to the request
     * @param {string} baseUrl The base url to append to
     * @param {object} params An object of query params. Will be properly formatted
     * @returns {string}
     */
    static addParams(baseUrl, params) {
        if (Object.keys(params).length === 0) {
            return baseUrl;
        }

        let paramToString = param => {
            if (Array.isArray(param)) {
                return param.map(paramToString).join(",");
            }

            if (typeof param === "object" && param !== null) {
                return Object.values(param).map(paramToString).join(",");
            }

            return String(param);
        };

        return baseUrl + "?" + Object.keys(params).map(k => k + "=" + escape(encodeURI(paramToString(params[k])))).join("&");
    }

    /**
     * Makes an ajax request, using the given options, or defaults when unset
     * @param {Object} options - The options for the request. See Ajax.defaultRequestOptions
     * @param {string} options.url - The url for the request
     * @param {?string} options.method - The http method to use
     * @param {?Object} options.queryParams - An object containing names and values of query params
     * @param {?Object} options.jsonData - An object containing arbitrary json data to post or put
     * @returns {Promise}
     */
    static request(options = {}) {
        options = Object.assign(Ajax.defaultRequestOptions, options);

        let url = options.url;

        if (typeof url !== "string" || url.length < 1) {
            return Promise.reject(new Error(`Illegal url provided: '${url}'`));
        }

        if (!/^((https?:\/\/)?[\w-]+(\.[\w-]+)*\.?(:\d+)?(\/\S*)?)/gi.test(url)) {
            return Promise.reject(new Error(`Invalid url provided: '${url}'`));
        }

        if (options.url.indexOf("?") === -1 && options.queryParams) {
            url = Ajax.addParams(url, options.queryParams);
        }

        let method = options.method;

        if (typeof ["GET", "POST", "PUT", "DELETE"].find(m => m === method) === "undefined") {
            return Promise.reject(new Error(`Invalid method provided: '${method}'`));
        }

        let jsonData = options.jsonData;

        if (typeof jsonData !== "object" && jsonData !== null) {
            return Promise.reject(new Error(`Illegal jsonData provided: '${jsonData}'`));
        }

        if (typeof jsonData !== "object" && Object.keys(jsonData).length > 0 && typeof ["POST", "PUT"].find(m => m === method) !== "undefined") {
            return Promise.reject(new Error("Invalid request: jsonData is only supported for POST and PUT requests"));
        }

        return new Promise((resolve, reject) => {
            let req = new XMLHttpRequest();

            req.open(method, url, true);

            req.onreadystatechange = () => {
                if (req.readyState === 4) {
                    try {
                        let response = JSON.parse(req.responseText);

                        if (req.status === 200) {
                            resolve(response);
                        } else {
                            if (response.error) {
                                reject(new Error(response.error));
                            } else {
                                reject(new Error(`An unknown error occured: (${req.responseText})`));
                            }
                        }
                    } catch (e) {
                        reject(new Error(`Failed to parse response (Code: ${req.status}): ${e.message}. Original response: ${req.responseText}`));
                    }
                }
            };

            if (jsonData !== null && (method === "POST" || method === "PUT")) {
                req.send(JSON.stringify(jsonData));
            } else {
                req.send();
            }
        });
    }
}

/**
 * Default request options for Ajax.request. If the default is undefined, then the field is required and will cause
 * an error if it is not provided.
 * @type {{url: string}}
 */
Ajax.defaultRequestOptions = {
    url: undefined,
    method: "GET",
    queryParams: {},
    jsonData: null
};

module.exports = Ajax;
