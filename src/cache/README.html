<p>title: postman
description:
authors: icy
categories: OP projects
created: 2022-12-06
updated: 2022-12-07
version: 0.0.18</p>

<h5>EN | <a href="README_RU.md">RU</a></h5>

<h1>Your friendly neighbour mailing list service</h1>

<ul>
<li>Written in pure GO</li>
<li>Celery client integration</li>
<li>Email templates with gohtml syntax supported</li>
<li>Check if email was read</li>
<li>All around nice guy letting everybody know about latest gossip in town</li>
</ul>

<h2>API methods</h2>

<h4>Create new mailing list</h4>

<p><code>POST /api/list</code> - response is an ID for new mailing list</p>

<h4>Add user to mailing list</h4>

<p><code>POST /api/list/{id}</code> - where <a href="id">id</a> is ID list&rsquo;s ID gotten from creation
  Request&rsquo;s body is a JSON struct with field names corresponding to user&rsquo;s information needed for Email template see <a href="#add-user">Add user</a></p>

<h4>Add multiple users</h4>

<p><code>POST /api/list/{id}/batch</code> - where <a href="id">id</a> is ID list&rsquo;s ID gotten from creation
  Request&rsquo;s body is a JSON array with structs where any individual struct is the same as with single user addition</p>

<h4>Get all the users in mailing list with specific ID</h4>

<p><code>GET /api/list/{id}</code> - response is an array of JSON structs of individual user&rsquo;s information in mailing list with specific ID</p>

<h4>Get read statistics from mailing list</h4>

<p><code>GET api/list/{id}/stat</code> - get information about who read email from mailing list
  Checkout implementation <a href="#read-statistics">Read statistics</a></p>

<h4>Add email template to mailing list</h4>

<p><code>POST /api/list/{id}/template</code> - request&rsquo;s body is a email template in text/html<br />
  <a href="https://pkg.go.dev/html/template">html/template</a> is used for email template implementation. This means that all syntax features of gohtml format
  are available, see <a href="#basic-template-for-a-mailing-list">Basic template for a mailing list</a></p>

<h2>Examples</h2>

<h4>Add user</h4>

<blockquote>
<p><code>POST /api/list/{id}</code></p>
<pre tabindex="0" style="color:#f8f8f2;background-color:#282a36;"><code><span style="display:flex;"><span>{
</span></span><span style="display:flex;"><span>&#34;email&#34;: &#34;moya@govorit.net&#34;,
</span></span><span style="display:flex;"><span>&#34;name&#34;: &#34;Alice&#34;
</span></span><span style="display:flex;"><span>&#34;bonuses&#34;: &#34;1200&#34;,
</span></span><span style="display:flex;"><span>&#34;link&#34;: &#34;season.shop/id=4535345432&#34;
</span></span><span style="display:flex;"><span>}
</span></span></code></pre></blockquote>

<h4>Basic template for a mailing list</h4>
<pre tabindex="0" style="color:#f8f8f2;background-color:#282a36;"><code><span style="display:flex;"><span>&lt;<span style="color:#ff79c6">html</span>&gt;
</span></span><span style="display:flex;"><span>&lt;<span style="color:#ff79c6">body</span>&gt;
</span></span><span style="display:flex;"><span>     &lt;<span style="color:#ff79c6">h1</span>&gt;👋 Hello {{ .name}}&lt;/<span style="color:#ff79c6">h1</span>&gt;
</span></span><span style="display:flex;"><span>    &lt;<span style="color:#ff79c6">p</span>&gt;You have {{ .bonuses}} bonusSeas 🐳&lt;/<span style="color:#ff79c6">p</span>&gt;
</span></span><span style="display:flex;"><span>     &lt;<span style="color:#ff79c6">p</span>&gt;Spend them wisely on:
</span></span><span style="display:flex;"><span>     &lt;<span style="color:#ff79c6">ul</span>&gt;
</span></span><span style="display:flex;"><span>     &lt;<span style="color:#ff79c6">li</span>&gt;Whale rental 🐋&lt;/<span style="color:#ff79c6">li</span>&gt;
</span></span><span style="display:flex;"><span>     &lt;<span style="color:#ff79c6">li</span>&gt;Steamboat tickets 💨&lt;/<span style="color:#ff79c6">li</span>&gt;
</span></span><span style="display:flex;"><span>     &lt;<span style="color:#ff79c6">li</span>&gt;Investing in sea 🪣&lt;/<span style="color:#ff79c6">li</span>&gt; 
</span></span><span style="display:flex;"><span>     &lt;/<span style="color:#ff79c6">ul</span>&gt;
</span></span><span style="display:flex;"><span>     &lt;/<span style="color:#ff79c6">p</span>&gt;
</span></span><span style="display:flex;"><span>     &lt;<span style="color:#ff79c6">p</span>&gt;More inforamtion in your profile &lt;<span style="color:#ff79c6">a</span> <span style="color:#50fa7b">href</span><span style="color:#ff79c6">=</span><span style="color:#f1fa8c">&#34;{{ .link}}&#34;</span>&gt;&lt;<span style="color:#ff79c6">span</span>&gt;here&lt;/<span style="color:#ff79c6">span</span>&gt; &lt;/<span style="color:#ff79c6">a</span>&gt;&lt;/<span style="color:#ff79c6">p</span>&gt;
</span></span><span style="display:flex;"><span>&lt;/<span style="color:#ff79c6">body</span>&gt;
</span></span><span style="display:flex;"><span>&lt;/<span style="color:#ff79c6">html</span>&gt;
</span></span><span style="display:flex;"><span>## There are multiple better ways to construct an email HTML template, this is just a simple example
</span></span></code></pre>
<h4>This example user information and email template produce the following email</h4>

<blockquote>
<p><strong>To:</strong> <em>moya@govorit.net</em><br />
<strong>From:</strong> <em>sea@son.shop</em>
<html>
  <body>
  <h1>👋 Hello Alice</h1>
  <p>You have 1200 bonusSeas 🐳</p>
  <p>Spend them wisely on:
  <ul>
  <li>Whale rental 🐋</li>
  <li>Steamboat tickets 💨</li>
  <li>Investing in sea 🪣</li>
  </ul>
  </p>
  <p>More inforamtion in your profile <a href="season.shop"><span>here</span> </a></p>
  </body>
  </html></p>
</blockquote>

<h2>Read statistics</h2>

<p>Checking if an email has been read is implemented using <a href="https://datatracker.ietf.org/doc/html/rfc3798">RFC3798</a> header<br />
You can also use service handler <code>GET /api/{email}/read</code> to deliberately register that user with this <code>{email}</code> has read the message</p>

<h2>Usage with celery</h2>

<p>Function <code>postman.mail(id)</code> takes a mailing list <a href="id">id</a> as an argument and starts constructing and sending emails based on template provided for mailing list with this ID</p>
