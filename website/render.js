class Page {
    id;
    attributes;
    elements;
}

/**
 * Render an object into page HTML.
 * @param p {object} an object to be rendered
 * @throws TypeError if p is formatted incorrectly
 * @return {string} HTML of the provided page
 */
function page(p) {
    let attrs = [];
    for (let dtype of attrOrder) {
        if (p.attributes[dtype] !== undefined) {
            attrs.push(renderAttr(p.attributes[dtype]));
        }
    }

    let elems = [];
    for (let e of p.elements.toSorted(
        (e1, e2) => e1.rendOrder < e2.rendOrder
    )) {
        elems.push(renderElement(e));
    }

    let attrsHtml = attrs.join('\n');
    let elemsHtml = elems.join('\n');

    return /* HTML */ `
        <div class="page">
            <div class="attributes">${attrsHtml}</div>
            <div class="elements">${elemsHtml}</div>
        </div>
    `;
}

// region attributes
class Attribute {
    /**
     * @type {string}
     */
    id;

    /**
     * @type {string}
     */
    css;

    /**
     * @type {string}
     */
    dtype;

    /**
     * @type {Title | object}
     */
    data;
}

class Title {
    /**
     * @type {string}
     */
    text;
}

const attrOrder = ['title'];

const renderers = {
    title: renderTitle,
};

/**
 * @param attr {Attribute}
 * @throws {TypeError}
 * @return {string}
 */
function renderAttr(attr) {
    if (renderers[attr.dtype] === undefined) {
        throw new TypeError(
            `cannot render attribute of unknown type: ${JSON.stringify(attr)}`
        );
    }

    return /* HTML */ `
        <div class="${attr.css}" id="${attr.id}">
            ${renderers[attr.dtype](attr.data)}
        </div>
    `;
}

/**
 * @param data {Title}
 * @throws {TypeError}
 * @return {string}
 */
function renderTitle(data) {
    return /* HTML */ ` <h1>${data.text}</h1> `;
}

// endregion

//region elements
class Element {
    id;
    rendOrder;
    css;
    ctype;
    content;
}

class Button {
    title;
    link;
}

const elementRenderers = {
    button: renderButton,
};

function renderElement(el) {
    if (elementRenderers[el.ctype] === undefined) {
        throw new TypeError(
            `tried to render an element of an unknown type: ${JSON.stringify(el)}`
        );
    }

    return /* HTML */ `
        <div class="${el.css}" id="${el.id}">
            ${elementRenderers[el.ctype](el.content)}
        </div>
    `;
}

function renderButton(content) {
    return /* HTML */ `
        <a href="${content.link}">
            <button>${content.title}</button>
        </a>
    `;
}

//endregion
