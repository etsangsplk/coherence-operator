<doc-view>

<v-layout row wrap>
<v-flex xs12 sm10 lg10>
<v-card class="section-def" v-bind:color="$store.state.currentColor">
<v-card-text class="pa-3">
<v-card class="section-def__card">
<v-card-text>
<dl>
<dt slot=title>Coherence Operator Development</dt>
<dd slot="desc"><p>The Coherence Operator is a Go based project built using the
<a id="" title="" target="_blank" href="https://github.com/operator-framework/operator-sdk">Operator SDK</a>.</p>
</dd>
</dl>
</v-card-text>
</v-card>
</v-card-text>
</v-card>
</v-flex>
</v-layout>

<h2 id="_coherence_operator_development">Coherence Operator Development</h2>
<div class="section">

<h3 id="_development_prerequisites">Development Prerequisites</h3>
<div class="section">
<p>The following prerequisites are required to build and test the operator (the prerequisites to just run the operator
are obviously a sub-set of these).</p>

<ul class="ulist">
<li>
<p><a id="" title="" target="_blank" href="https://github.com/operator-framework/operator-sdk/tree/v0.13.0">operator-sdk</a> version <strong>v0.13.0</strong></p>

</li>
<li>
<p><a id="" title="" target="_blank" href="https://git-scm.com/downloads">git</a></p>

</li>
<li>
<p><a id="" title="" target="_blank" href="https://golang.org/dl/">go</a> version v1.13.4+.</p>

</li>
<li>
<p><a id="" title="" target="_blank" href="https://www.mercurial-scm.org/downloads">mercurial</a> version 3.9+</p>

</li>
<li>
<p><a id="" title="" target="_blank" href="https://docs.docker.com/install/">docker</a> version 17.03+.</p>

</li>
<li>
<p><a id="" title="" target="_blank" href="https://kubernetes.io/docs/tasks/tools/install-kubectl/">kubectl</a> version v1.12.0+.</p>

</li>
<li>
<p>Access to a Kubernetes v1.12.0+ cluster.</p>

</li>
<li>
<p><a id="" title="" target="_blank" href="http://jdk.java.net/">Java 8+ JDK</a></p>

</li>
<li>
<p><a id="" title="" target="_blank" href="https://maven.apache.org">Maven</a> version 3.5+</p>

</li>
<li>
<p>Access to a Maven repository containing Oracle Coherence 12.2.1.4 (for the exact GAV see the
<code>pom.xml</code> file in the <code>java/</code> directory)</p>

</li>
<li>
<p>Optional: <a id="" title="" target="_blank" href="https://github.com/go-delve/delve/tree/master/Documentation/installation">delve</a>
version 1.2.0+ (for local debugging with <code>operator-sdk up local --enable-delve</code>).</p>

</li>
<li>
<p>This project uses <code>make</code> for building, which should already be installed on most systems</p>

</li>
</ul>
<div class="admonition note">
<p class="admonition-inline">This project currently uses the Operator SDK v0.11.0 so make sure you install the correct version of
the Operator SDK CLI.</p>
</div>
<div class="admonition note">
<p class="admonition-inline">As stated above this project requires K8s v1.12.0+ so if using Docker on MacOS you need at least version 2.1.0.0</p>
</div>
</div>

<h3 id="_project_structure">Project Structure</h3>
<div class="section">
<p>The project also contains a Java sub-project that is used to create Coherence utilities that the Operator relies on
to work correctly with Coherence clusters that it is managing.</p>

<p>This project was initially generated using the Operator SDK and this dictates the structure of the project
which means that files and directories should not be moved arbitrarily.</p>

</div>

<h3 id="_operator_sdk_files">Operator SDK Files</h3>
<div class="section">
<p>The following should not be moved:</p>


<div class="table__overflow elevation-1 ">
<table class="datatable table">
<colgroup>
<col style="width: 50%;">
<col style="width: 50%;">
</colgroup>
<thead>
<tr>
<th>File</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>bin/</code></td>
<td>scripts used in the Operator Docker image</td>
</tr>
<tr>
<td><code>build/Dockerfile</code></td>
<td>the <code>Dockerfile</code> used by the Operator SDK to build the Docker image</td>
</tr>
<tr>
<td><code>cmd/manager/main.go</code></td>
<td>The Operator <code>main</code> generated by the Operator SDK</td>
</tr>
<tr>
<td><code>deploy/</code></td>
<td>Yaml files generated and maintained by the Operator SDK</td>
</tr>
<tr>
<td><code>deploy/crds</code></td>
<td>The CRD files generated and maintained by the Operator SDK</td>
</tr>
<tr>
<td><code>helm-charts/</code></td>
<td>The Helm charts used by the Operator</td>
</tr>
<tr>
<td><code>pkg/apis</code></td>
<td>The API <code>struct</code> code generated by the Operator SDK and used to generate the CRD files</td>
</tr>
<tr>
<td><code>pkg/controller</code></td>
<td>The controller code generated by the Operator SDK</td>
</tr>
<tr>
<td><code>watches.yaml</code></td>
<td>The Helm Operator configuration generated by the Operator SDK</td>
</tr>
<tr>
<td><code>local-watches.yaml</code></td>
<td>The Helm Operator configuration used when running the operator locally</td>
</tr>
</tbody>
</table>
</div>
</div>
</div>
</doc-view>