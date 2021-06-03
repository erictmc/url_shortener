import camelCase from "lodash.camelcase";
import isObject from "lodash.isobject";
import snakeCase from "lodash.snakecase";

// https://stackoverflow.com/questions/3809401/what-is-a-good-regular-expression-to-match-a-url
// Raw regex is used, so we can test consistently client and server side
const urlRegex = '^(?:(?:http|https)://)(?:\\S+(?::\\S*)?@)?(?:(?:(?:[1-9]\\d?|1\\d\\d|2[01]\\d|22[0-3])(?:\\.(?:1?\\d{1,2}|2[0-4]\\d|25[0-5])){2}(?:\\.(?:[0-9]\\d?|1\\d\\d|2[0-4]\\d|25[0-4]))|(?:(?:[a-z\\u00a1-\\uffff0-9]+-?)*[a-z\\u00a1-\\uffff0-9]+)(?:\\.(?:[a-z\\u00a1-\\uffff0-9]+-?)*[a-z\\u00a1-\\uffff0-9]+)*(?:\\.(?:[a-z\\u00a1-\\uffff]{2,})))|localhost)(?::\\d{2,5})?(?:(/|\\?|#)[^\\s]*)?$'

export const isValidUrl = (url: string) => {
    return url.match(urlRegex) !== null
}

export const convertObjToCamelCase = (obj: object|object[]): object|object[] => {
    return applyFuncToObjectFields(obj, camelCase)
};

export const convertObjToSnakeCase = (obj: object|object[]): object|object[] => {
    return applyFuncToObjectFields(obj, snakeCase)
};

// Helper function for changing cases of object fields
// @ts-ignore
const applyFuncToObjectFields = (o, func: Function)  => {
    let newO, origKey, newKey, value;
    if (o instanceof Array) {
        return o.map(function(value) {
            if (isObject(value)) {
                value = applyFuncToObjectFields(value, func);
            }
            return value;
        });
    } else {
        newO = {};
        for (origKey in o) {
            if (o.hasOwnProperty(origKey)) {
                newKey = func(origKey);
                value = o[origKey];
                if (
                    value instanceof Array ||
                    (value !== null && value !== undefined && value.constructor === Object)
                ) {
                    value = applyFuncToObjectFields(value, func);
                }
                // @ts-ignore
                newO[newKey] = value;
            }
        }
    }
    return newO;
};