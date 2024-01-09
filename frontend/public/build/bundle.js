var app=function(){"use strict";function t(){}function e(t){return t()}function n(){return Object.create(null)}function l(t){t.forEach(e)}function r(t){return"function"==typeof t}function c(t,e){return t!=t?e==e:t!==e||t&&"object"==typeof t||"function"==typeof t}function o(t,e){t.appendChild(e)}function i(t,e,n){t.insertBefore(e,n||null)}function u(t){t.parentNode&&t.parentNode.removeChild(t)}function a(t,e){for(let n=0;n<t.length;n+=1)t[n]&&t[n].d(e)}function s(t){return document.createElement(t)}function d(t){return document.createTextNode(t)}function f(){return d(" ")}function h(){return d("")}function p(t,e,n,l){return t.addEventListener(e,n,l),()=>t.removeEventListener(e,n,l)}function m(t,e,n){null==n?t.removeAttribute(e):t.getAttribute(e)!==n&&t.setAttribute(e,n)}function g(t,e){e=""+e,t.data!==e&&(t.data=e)}function $(t,e){t.value=null==e?"":e}let b;function v(t){b=t}function x(){if(!b)throw new Error("Function called outside component initialization");return b}function y(t){x().$$.after_update.push(t)}const k=[],w=[];let C=[];const E=[],_=Promise.resolve();let D=!1;function H(t){C.push(t)}const I=new Set;let N=0;function A(){if(0!==N)return;const t=b;do{try{for(;N<k.length;){const t=k[N];N++,v(t),T(t.$$)}}catch(t){throw k.length=0,N=0,t}for(v(null),k.length=0,N=0;w.length;)w.pop()();for(let t=0;t<C.length;t+=1){const e=C[t];I.has(e)||(I.add(e),e())}C.length=0}while(k.length);for(;E.length;)E.pop()();D=!1,I.clear(),v(t)}function T(t){if(null!==t.fragment){t.update(),l(t.before_update);const e=t.dirty;t.dirty=[-1],t.fragment&&t.fragment.p(t.ctx,e),t.after_update.forEach(H)}}const q=new Set;let S;function L(){S={r:0,c:[],p:S}}function R(){S.r||l(S.c),S=S.p}function O(t,e){t&&t.i&&(q.delete(t),t.i(e))}function j(t,e,n,l){if(t&&t.o){if(q.has(t))return;q.add(t),S.c.push((()=>{q.delete(t),l&&(n&&t.d(1),l())})),t.o(e)}else l&&l()}function X(t,e){const n=e.token={};function l(t,l,r,c){if(e.token!==n)return;e.resolved=c;let o=e.ctx;void 0!==r&&(o=o.slice(),o[r]=c);const i=t&&(e.current=t)(o);let u=!1;e.block&&(e.blocks?e.blocks.forEach(((t,n)=>{n!==l&&t&&(L(),j(t,1,1,(()=>{e.blocks[n]===t&&(e.blocks[n]=null)})),R())})):e.block.d(1),i.c(),O(i,1),i.m(e.mount(),e.anchor),u=!0),e.block=i,e.blocks&&(e.blocks[l]=i),u&&A()}if(!(r=t)||"object"!=typeof r&&"function"!=typeof r||"function"!=typeof r.then){if(e.current!==e.then)return l(e.then,1,e.value,t),!0;e.resolved=t}else{const n=x();if(t.then((t=>{v(n),l(e.then,1,e.value,t),v(null)}),(t=>{if(v(n),l(e.catch,2,e.error,t),v(null),!e.hasCatch)throw t})),e.current!==e.pending)return l(e.pending,0),!0}var r}function P(t){t&&t.c()}function B(t,n,c,o){const{fragment:i,after_update:u}=t.$$;i&&i.m(n,c),o||H((()=>{const n=t.$$.on_mount.map(e).filter(r);t.$$.on_destroy?t.$$.on_destroy.push(...n):l(n),t.$$.on_mount=[]})),u.forEach(H)}function F(t,e){const n=t.$$;null!==n.fragment&&(!function(t){const e=[],n=[];C.forEach((l=>-1===t.indexOf(l)?e.push(l):n.push(l))),n.forEach((t=>t())),C=e}(n.after_update),l(n.on_destroy),n.fragment&&n.fragment.d(e),n.on_destroy=n.fragment=null,n.ctx=[])}function Q(t,e){-1===t.$$.dirty[0]&&(k.push(t),D||(D=!0,_.then(A)),t.$$.dirty.fill(0)),t.$$.dirty[e/31|0]|=1<<e%31}function z(e,r,c,o,i,a,s,d=[-1]){const f=b;v(e);const h=e.$$={fragment:null,ctx:[],props:a,update:t,not_equal:i,bound:n(),on_mount:[],on_destroy:[],on_disconnect:[],before_update:[],after_update:[],context:new Map(r.context||(f?f.$$.context:[])),callbacks:n(),dirty:d,skip_bound:!1,root:r.target||f.$$.root};s&&s(h.root);let p=!1;if(h.ctx=c?c(e,r.props||{},((t,n,...l)=>{const r=l.length?l[0]:n;return h.ctx&&i(h.ctx[t],h.ctx[t]=r)&&(!h.skip_bound&&h.bound[t]&&h.bound[t](r),p&&Q(e,t)),n})):[],h.update(),p=!0,l(h.before_update),h.fragment=!!o&&o(h.ctx),r.target){if(r.hydrate){const t=function(t){return Array.from(t.childNodes)}(r.target);h.fragment&&h.fragment.l(t),t.forEach(u)}else h.fragment&&h.fragment.c();r.intro&&O(e.$$.fragment),B(e,r.target,r.anchor,r.customElement),A()}v(f)}class G{$destroy(){F(this,1),this.$destroy=t}$on(e,n){if(!r(n))return t;const l=this.$$.callbacks[e]||(this.$$.callbacks[e]=[]);return l.push(n),()=>{const t=l.indexOf(n);-1!==t&&l.splice(t,1)}}$set(t){var e;this.$$set&&(e=t,0!==Object.keys(e).length)&&(this.$$.skip_bound=!0,this.$$set(t),this.$$.skip_bound=!1)}}const J=["REAL","TEXT"],M=[0,""],K={REAL:0,TEXT:1};function U(t,e,n){const l=t.slice();return l[11]=e[n],l[13]=n,l}function V(t,e,n){const l=t.slice();return l[14]=e[n],l[15]=e,l[16]=n,l}function W(t,e,n){const l=t.slice();return l[17]=e[n],l}function Y(t){let e,n,l,r,c,a,h=t[17].name+"",p=J[t[17].typeID]+"";return{c(){e=s("div"),n=s("div"),l=d(h),r=f(),c=s("div"),a=d(p),m(n,"class","centered-text col-name"),m(c,"class","centered-text col-type"),m(e,"class","flexible")},m(t,u){i(t,e,u),o(e,n),o(n,l),o(e,r),o(e,c),o(c,a)},p(t,e){2&e&&h!==(h=t[17].name+"")&&g(l,h),2&e&&p!==(p=J[t[17].typeID]+"")&&g(a,p)},d(t){t&&u(e)}}}function Z(t){let e,n,l,r;function c(){t[7].call(n,t[15],t[16])}return{c(){e=s("div"),n=s("input"),m(e,"class","tall flexible")},m(u,a){i(u,e,a),o(e,n),$(n,t[14]),l||(r=p(n,"input",c),l=!0)},p(e,l){t=e,4&l&&n.value!==t[14]&&$(n,t[14])},d(t){t&&u(e),l=!1,r()}}}function tt(t){let e,n,l,r,c,d=t[11],h=[];for(let e=0;e<d.length;e+=1)h[e]=Z(V(t,d,e));function g(){return t[8](t[13])}return{c(){e=s("div");for(let t=0;t<h.length;t+=1)h[t].c();n=f(),l=s("button"),l.textContent="X",m(l,"class","unflexible tall padded-element exit-button"),m(e,"class","flex row-height divider")},m(t,u){i(t,e,u);for(let t=0;t<h.length;t+=1)h[t]&&h[t].m(e,null);o(e,n),o(e,l),r||(c=p(l,"click",g),r=!0)},p(l,r){if(t=l,4&r){let l;for(d=t[11],l=0;l<d.length;l+=1){const c=V(t,d,l);h[l]?h[l].p(c,r):(h[l]=Z(c),h[l].c(),h[l].m(e,n))}for(;l<h.length;l+=1)h[l].d(1);h.length=d.length}},d(t){t&&u(e),a(h,t),r=!1,c()}}}function et(e){let n,r,c,h,$,b,v,x,y,k,w,C,E,_,D,H=e[1],I=[];for(let t=0;t<H.length;t+=1)I[t]=Y(W(e,H,t));let N=e[2],A=[];for(let t=0;t<N.length;t+=1)A[t]=tt(U(e,N,t));return{c(){n=s("div"),r=s("div"),c=d(e[0]),h=f(),$=s("div");for(let t=0;t<I.length;t+=1)I[t].c();b=f(),v=s("div"),v.textContent="X",x=f();for(let t=0;t<A.length;t+=1)A[t].c();y=f(),k=s("div"),w=s("button"),w.textContent="Add Row",C=f(),E=s("button"),E.textContent="Submit Request",m(r,"class","title-bar"),m(v,"class","unflexible padded-element placeholder"),m($,"class","flex divider"),m(w,"class","flexible padded-element simple-button"),m(E,"class","flexible padded-element simple-button"),m(k,"class","flex")},m(t,l){i(t,n,l),o(n,r),o(r,c),o(n,h),o(n,$);for(let t=0;t<I.length;t+=1)I[t]&&I[t].m($,null);o($,b),o($,v),o(n,x);for(let t=0;t<A.length;t+=1)A[t]&&A[t].m(n,null);o(n,y),o(n,k),o(k,w),o(k,C),o(k,E),_||(D=[p(w,"click",e[3]),p(E,"click",e[5])],_=!0)},p(t,[e]){if(1&e&&g(c,t[0]),2&e){let n;for(H=t[1],n=0;n<H.length;n+=1){const l=W(t,H,n);I[n]?I[n].p(l,e):(I[n]=Y(l),I[n].c(),I[n].m($,b))}for(;n<I.length;n+=1)I[n].d(1);I.length=H.length}if(20&e){let l;for(N=t[2],l=0;l<N.length;l+=1){const r=U(t,N,l);A[l]?A[l].p(r,e):(A[l]=tt(r),A[l].c(),A[l].m(n,y))}for(;l<A.length;l+=1)A[l].d(1);A.length=N.length}},i:t,o:t,d(t){t&&u(n),a(I,t),a(A,t),_=!1,l(D)}}}function nt(t,e,n){let{tableName:l}=e,{structure:r}=e,{dataHandler:c}=e,o=[],i=l,u=r;function a(t){o.splice(t,1),n(2,o)}y((()=>{l==i&&r==u||(n(2,o=[]),i=l,u=r)}));return t.$$set=t=>{"tableName"in t&&n(0,l=t.tableName),"structure"in t&&n(1,r=t.structure),"dataHandler"in t&&n(6,c=t.dataHandler)},[l,r,o,function(){let t=r.map((t=>M[t.typeID]));o.push(t),n(2,o)},a,function(){if(0==o.length)return void alert("Cannot submit empty request.");let t=new Array(o.length);for(let e=0;e<o.length;e++){let n={};for(let t=0;t<r.length;t++){if(r[t].typeID==K.REAL){let l=parseFloat(o[e][t]);if(isNaN(l))return void alert("REAL required at ("+e+", "+t+")");n[r[t].name]=l}else n[r[t].name]=o[e][t]}t[e]=n}c(l,t)},c,function(t,e){t[e]=this.value,n(2,o)},t=>a(t)]}class lt extends G{constructor(t){super(),z(this,t,nt,et,c,{tableName:0,structure:1,dataHandler:6})}}function rt(t,e,n){const l=t.slice();return l[6]=e[n],l[8]=n,l}function ct(t){let e,n,l,r,c,a=t[6]+"";function h(){return t[5](t[8])}return{c(){e=s("button"),n=d(a),l=f(),m(e,"class","tall padded-element untoggled")},m(t,u){i(t,e,u),o(e,n),o(e,l),r||(c=p(e,"click",h),r=!0)},p(e,l){t=e,2&l&&a!==(a=t[6]+"")&&g(n,a)},d(t){t&&u(e),r=!1,c()}}}function ot(t){let e,n,l,r,c,a=t[6]+"";function h(){return t[4](t[8])}return{c(){e=s("button"),n=d(a),l=f(),m(e,"class","tall padded-element toggled")},m(t,u){i(t,e,u),o(e,n),o(e,l),r||(c=p(e,"click",h),r=!0)},p(e,l){t=e,2&l&&a!==(a=t[6]+"")&&g(n,a)},d(t){t&&u(e),r=!1,c()}}}function it(t){let e;function n(t,e){return t[8]==t[0]?ot:ct}let l=n(t),r=l(t);return{c(){r.c(),e=h()},m(t,n){r.m(t,n),i(t,e,n)},p(t,c){l===(l=n(t))&&r?r.p(t,c):(r.d(1),r=l(t),r&&(r.c(),r.m(e.parentNode,e)))},d(t){r.d(t),t&&u(e)}}}function ut(e){let n,l=e[1],r=[];for(let t=0;t<l.length;t+=1)r[t]=it(rt(e,l,t));return{c(){n=s("div");for(let t=0;t<r.length;t+=1)r[t].c();m(n,"class","tall")},m(t,e){i(t,n,e);for(let t=0;t<r.length;t+=1)r[t]&&r[t].m(n,null)},p(t,[e]){if(7&e){let c;for(l=t[1],c=0;c<l.length;c+=1){const o=rt(t,l,c);r[c]?r[c].p(o,e):(r[c]=it(o),r[c].c(),r[c].m(n,null))}for(;c<r.length;c+=1)r[c].d(1);r.length=l.length}},i:t,o:t,d(t){t&&u(n),a(r,t)}}}function at(t,e,n){let{choices:l}=e,{ci:r}=e,{choiceHandler:c}=e;function o(t){c(l[t]),n(0,r=t)}return t.$$set=t=>{"choices"in t&&n(1,l=t.choices),"ci"in t&&n(0,r=t.ci),"choiceHandler"in t&&n(3,c=t.choiceHandler)},[r,l,o,c,t=>o(t),t=>o(t)]}class st extends G{constructor(t){super(),z(this,t,at,ut,c,{choices:1,ci:0,choiceHandler:3})}}function dt(t,e,n){const l=t.slice();return l[10]=e[n],l[11]=e,l[12]=n,l}function ft(t){let e,n,r,c,a,d,h,g,b,v,x,y,k;function w(){t[7].call(r,t[11],t[12])}function C(...e){return t[8](t[12],...e)}function E(){return t[9](t[12])}return d=new st({props:{choices:J,ci:t[1][t[12]].typeID,choiceHandler:C}}),{c(){e=s("div"),n=s("div"),r=s("input"),c=f(),a=s("div"),P(d.$$.fragment),h=f(),g=s("div"),b=s("button"),b.textContent="X",v=f(),m(r,"placeholder","Column"),m(n,"class","flexible"),m(a,"class","inflexible"),m(b,"class","tall padded-element exit-button"),m(g,"class","inflexible"),m(e,"class","struct-row divider svelte-fsd72d")},m(l,u){i(l,e,u),o(e,n),o(n,r),$(r,t[10].name),o(e,c),o(e,a),B(d,a,null),o(e,h),o(e,g),o(g,b),o(e,v),x=!0,y||(k=[p(r,"input",w),p(b,"click",E)],y=!0)},p(e,n){t=e,2&n&&r.value!==t[10].name&&$(r,t[10].name);const l={};2&n&&(l.ci=t[1][t[12]].typeID),2&n&&(l.choiceHandler=C),d.$set(l)},i(t){x||(O(d.$$.fragment,t),x=!0)},o(t){j(d.$$.fragment,t),x=!1},d(t){t&&u(e),F(d),y=!1,l(k)}}}function ht(t){let e,n,r,c,d,h,g,b,v,x,y,k,w,C=t[1],E=[];for(let e=0;e<C.length;e+=1)E[e]=ft(dt(t,C,e));const _=t=>j(E[t],1,1,(()=>{E[t]=null}));return{c(){e=s("div"),n=s("div"),r=s("div"),c=s("input"),d=f();for(let t=0;t<E.length;t+=1)E[t].c();h=f(),g=s("div"),b=s("button"),b.textContent="Add Column",v=f(),x=s("button"),x.textContent="Create Request",m(c,"placeholder","Table"),m(r,"class","struct-row divider svelte-fsd72d"),m(n,"class","struct-table"),m(b,"class","flexible padded-element simple-button"),m(x,"class","flexible padded-element simple-button"),m(g,"class","flex"),m(e,"class","struct-container wide")},m(l,u){i(l,e,u),o(e,n),o(n,r),o(r,c),$(c,t[0]),o(n,d);for(let t=0;t<E.length;t+=1)E[t]&&E[t].m(n,null);o(e,h),o(e,g),o(g,b),o(g,v),o(g,x),y=!0,k||(w=[p(c,"input",t[6]),p(b,"click",t[2]),p(x,"click",t[4])],k=!0)},p(t,[e]){if(1&e&&c.value!==t[0]&&$(c,t[0]),10&e){let l;for(C=t[1],l=0;l<C.length;l+=1){const r=dt(t,C,l);E[l]?(E[l].p(r,e),O(E[l],1)):(E[l]=ft(r),E[l].c(),O(E[l],1),E[l].m(n,null))}for(L(),l=C.length;l<E.length;l+=1)_(l);R()}},i(t){if(!y){for(let t=0;t<C.length;t+=1)O(E[t]);y=!0}},o(t){E=E.filter(Boolean);for(let t=0;t<E.length;t+=1)j(E[t]);y=!1},d(t){t&&u(e),a(E,t),k=!1,l(w)}}}function pt(t,e,n){let{structureHandler:l}=e,r="",c=[];function o(t){c.splice(t,1),n(1,c)}return t.$$set=t=>{"structureHandler"in t&&n(5,l=t.structureHandler)},[r,c,function(){c.push({name:"",typeID:0}),n(1,c)},o,function(){if(0==r.length)return void alert("Please enter a table name.");if(0==c.length)return void alert("Please add columns.");let t=new Set;for(const e of c){if(0==e.name.length)return void alert("Column names must be non-empty.");if(t.has(e.name))return void alert("Column names must be unique.");t.add(e.name)}l(r,c.map((t=>({name:t.name,typeID:t.typeID}))))},l,function(){r=this.value,n(0,r)},function(t,e){t[e].name=this.value,n(1,c)},(t,e)=>n(1,c[t].typeID=K[e],c),t=>o(t)]}class mt extends G{constructor(t){super(),z(this,t,pt,ht,c,{structureHandler:5})}}const gt="134.122.3.34:3000",$t="http://"+gt+"/query",bt="http://"+gt+"/data";function vt(t){let e,n,l,r,c;return r=new lt({props:{tableName:t[0],structure:t[1],dataHandler:t[3]}}),{c(){e=s("div"),n=f(),l=s("div"),P(r.$$.fragment),m(e,"class","break"),m(l,"class","rounded")},m(t,o){i(t,e,o),i(t,n,o),i(t,l,o),B(r,l,null),c=!0},p(t,e){const n={};1&e&&(n.tableName=t[0]),2&e&&(n.structure=t[1]),r.$set(n)},i(t){c||(O(r.$$.fragment,t),c=!0)},o(t){j(r.$$.fragment,t),c=!1},d(t){t&&u(e),t&&u(n),t&&u(l),F(r)}}}function xt(t){let e,n,l,r,c,a,d,h;a=new mt({props:{structureHandler:t[2]}});let p=t[1].length>0&&vt(t);return{c(){e=s("div"),e.textContent="Insert Data",n=f(),l=s("div"),r=s("div"),c=s("div"),P(a.$$.fragment),d=f(),p&&p.c(),m(e,"class","title-bar"),m(c,"class","rounded"),m(r,"class","tight"),m(l,"class","container")},m(t,u){i(t,e,u),i(t,n,u),i(t,l,u),o(l,r),o(r,c),B(a,c,null),o(l,d),p&&p.m(l,null),h=!0},p(t,[e]){t[1].length>0?p?(p.p(t,e),2&e&&O(p,1)):(p=vt(t),p.c(),O(p,1),p.m(l,null)):p&&(L(),j(p,1,1,(()=>{p=null})),R())},i(t){h||(O(a.$$.fragment,t),O(p),h=!0)},o(t){j(a.$$.fragment,t),j(p),h=!1},d(t){t&&u(e),t&&u(n),t&&u(l),F(a),p&&p.d()}}}function yt(t,e,n){let l="",r=[];return[l,r,function(t,e){n(0,l=t),n(1,r=e)},async function(t,e){console.log(e);let n=await fetch(bt+"?table="+t,{method:"POST",headers:{"Content-Type":"application/x-www-form-urlencoded"},body:JSON.stringify(e)}),l=await n.json();n.ok?alert("Insert Success!"):alert(l.message)}]}class kt extends G{constructor(t){super(),z(this,t,yt,xt,c,{})}}function wt(t,e,n){const l=t.slice();return l[7]=e[n],l}function Ct(t,e,n){const l=t.slice();return l[10]=e[n],l[11]=e,l[12]=n,l}function Et(t,e,n){const l=t.slice();return l[13]=e[n],l}function _t(t){let e,n,l,r,c,a,h,p=t[13].name+"",$=J[t[13].typeID]+"";return{c(){e=s("div"),n=s("div"),l=d(p),r=f(),c=s("div"),a=d($),h=f(),m(n,"class","centered-text col-name"),m(c,"class","centered-text col-type"),m(e,"class","flexible")},m(t,u){i(t,e,u),o(e,n),o(n,l),o(e,r),o(e,c),o(c,a),o(e,h)},p(t,e){1&e&&p!==(p=t[13].name+"")&&g(l,p),1&e&&$!==($=J[t[13].typeID]+"")&&g(a,$)},d(t){t&&u(e)}}}function Dt(t){let e,n,l,r;function c(){t[3].call(n,t[11],t[12])}return{c(){e=s("div"),n=s("input"),n.readOnly=!0,m(e,"class","tall flexible")},m(u,a){i(u,e,a),o(e,n),$(n,t[10]),l||(r=p(n,"input",c),l=!0)},p(e,l){t=e,2&l&&n.value!==t[10]&&$(n,t[10])},d(t){t&&u(e),l=!1,r()}}}function Ht(t){let e,n,l=t[7],r=[];for(let e=0;e<l.length;e+=1)r[e]=Dt(Ct(t,l,e));return{c(){e=s("div");for(let t=0;t<r.length;t+=1)r[t].c();n=f(),m(e,"class","flex row-height divider")},m(t,l){i(t,e,l);for(let t=0;t<r.length;t+=1)r[t]&&r[t].m(e,null);o(e,n)},p(t,c){if(2&c){let o;for(l=t[7],o=0;o<l.length;o+=1){const i=Ct(t,l,o);r[o]?r[o].p(i,c):(r[o]=Dt(i),r[o].c(),r[o].m(e,n))}for(;o<r.length;o+=1)r[o].d(1);r.length=l.length}},d(t){t&&u(e),a(r,t)}}}function It(e){let n,l,r,c=e[0],d=[];for(let t=0;t<c.length;t+=1)d[t]=_t(Et(e,c,t));let h=e[1],p=[];for(let t=0;t<h.length;t+=1)p[t]=Ht(wt(e,h,t));return{c(){n=s("div"),l=s("div");for(let t=0;t<d.length;t+=1)d[t].c();r=f();for(let t=0;t<p.length;t+=1)p[t].c();m(l,"class","flex divider")},m(t,e){i(t,n,e),o(n,l);for(let t=0;t<d.length;t+=1)d[t]&&d[t].m(l,null);o(n,r);for(let t=0;t<p.length;t+=1)p[t]&&p[t].m(n,null)},p(t,[e]){if(1&e){let n;for(c=t[0],n=0;n<c.length;n+=1){const r=Et(t,c,n);d[n]?d[n].p(r,e):(d[n]=_t(r),d[n].c(),d[n].m(l,null))}for(;n<d.length;n+=1)d[n].d(1);d.length=c.length}if(2&e){let l;for(h=t[1],l=0;l<h.length;l+=1){const r=wt(t,h,l);p[l]?p[l].p(r,e):(p[l]=Ht(r),p[l].c(),p[l].m(n,null))}for(;l<p.length;l+=1)p[l].d(1);p.length=h.length}},i:t,o:t,d(t){t&&u(n),a(d,t),a(p,t)}}}function Nt(t,e,n){let{data:l}=e,r=null,c=[],o=[];return y((()=>{r!=l&&(r=l,n(0,c=function(){let t=[],e=l[0];for(const n in e){let l=e[n],r=K.REAL;"string"==typeof l&&(r=K.TEXT),t.push({name:n,typeID:r})}return t}()),n(1,o=l.map((t=>c.map((e=>t[e.name]))))))})),t.$$set=t=>{"data"in t&&n(2,l=t.data)},[c,o,l,function(t,e){t[e]=this.value,n(1,o)}]}class At extends G{constructor(t){super(),z(this,t,Nt,It,c,{data:2})}}function Tt(e){return{c:t,m:t,p:t,i:t,o:t,d:t}}function qt(t){let e,n,l=null!=t[5]&&t[5].length>0&&St(t);return{c(){l&&l.c(),e=h()},m(t,r){l&&l.m(t,r),i(t,e,r),n=!0},p(t,n){null!=t[5]&&t[5].length>0?l?(l.p(t,n),2&n&&O(l,1)):(l=St(t),l.c(),O(l,1),l.m(e.parentNode,e)):l&&(L(),j(l,1,1,(()=>{l=null})),R())},i(t){n||(O(l),n=!0)},o(t){j(l),n=!1},d(t){l&&l.d(t),t&&u(e)}}}function St(t){let e,n,l,r;return l=new At({props:{data:t[5]}}),{c(){e=s("div"),n=f(),P(l.$$.fragment),m(e,"class","break")},m(t,c){i(t,e,c),i(t,n,c),B(l,t,c),r=!0},p(t,e){const n={};2&e&&(n.data=t[5]),l.$set(n)},i(t){r||(O(l.$$.fragment,t),r=!0)},o(t){j(l.$$.fragment,t),r=!1},d(t){t&&u(e),t&&u(n),F(l,t)}}}function Lt(e){let n;return{c(){n=s("div")},m(t,e){i(t,n,e)},p:t,i:t,o:t,d(t){t&&u(n)}}}function Rt(t){let e,n,r,c,a,d,h,g,b,v,x,y,k,w,C={ctx:t,current:null,token:null,hasCatch:!1,pending:Lt,then:qt,catch:Tt,value:5,blocks:[,,,]};return X(x=t[1],C),{c(){e=s("div"),n=s("div"),n.textContent="Query Data",r=f(),c=s("div"),a=s("div"),d=s("div"),h=s("textarea"),g=f(),b=s("button"),b.textContent="Submit",v=f(),C.block.c(),m(n,"class","title-bar"),m(h,"class","flexible padded-element"),m(h,"placeholder","Enter Query..."),m(d,"class","flex divider"),m(b,"class","simple-button padded-element wide"),m(a,"class","rounded"),m(c,"class","container")},m(l,u){i(l,e,u),o(e,n),o(e,r),o(e,c),o(c,a),o(a,d),o(d,h),$(h,t[0]),o(a,g),o(a,b),o(c,v),C.block.m(c,C.anchor=null),C.mount=()=>c,C.anchor=null,y=!0,k||(w=[p(h,"input",t[3]),p(b,"click",t[2])],k=!0)},p(e,[n]){t=e,1&n&&$(h,t[0]),C.ctx=t,2&n&&x!==(x=t[1])&&X(x,C)||function(t,e,n){const l=e.slice(),{resolved:r}=t;t.current===t.then&&(l[t.value]=r),t.current===t.catch&&(l[t.error]=r),t.block.p(l,n)}(C,t,n)},i(t){y||(O(C.block),y=!0)},o(t){for(let t=0;t<3;t+=1){j(C.blocks[t])}y=!1},d(t){t&&u(e),C.block.d(),C.token=null,C=null,k=!1,l(w)}}}function Ot(t,e,n){let l="",r=null;return[l,r,function(){0!=l.length?n(1,r=async function(){let t=await fetch($t+"?q="+l,{method:"GET"}),e=await t.json();return t.ok?e.data:(alert(e.message),[])}()):alert("Please enter a query.")},function(){l=this.value,n(0,l)}]}class jt extends G{constructor(t){super(),z(this,t,Ot,Rt,c,{})}}function Xt(e){let n,l,r,c,a,d,h,p,g,$,b,v;return d=new jt({}),b=new kt({}),{c(){n=s("main"),l=s("div"),l.textContent="ScratchDB Lite",r=f(),c=s("div"),a=s("div"),P(d.$$.fragment),h=f(),p=s("div"),g=f(),$=s("div"),P(b.$$.fragment),m(l,"class","title-bar"),m(a,"class","rounded section"),m(p,"class","break"),m($,"class","rounded section"),m(c,"class","container")},m(t,e){i(t,n,e),o(n,l),o(n,r),o(n,c),o(c,a),B(d,a,null),o(c,h),o(c,p),o(c,g),o(c,$),B(b,$,null),v=!0},p:t,i(t){v||(O(d.$$.fragment,t),O(b.$$.fragment,t),v=!0)},o(t){j(d.$$.fragment,t),j(b.$$.fragment,t),v=!1},d(t){t&&u(n),F(d),F(b)}}}return new class extends G{constructor(t){super(),z(this,t,null,Xt,c,{})}}({target:document.body})}();
//# sourceMappingURL=bundle.js.map
