<h1>Image Watermarking Tool</h1>

<p>This Go program adds a watermark to images in a specified folder. It takes user input for the folder path and creates watermarked images in a new folder.</p>

<h2>Features</h2>

<ul>
  <li>Watermarks images with a specified logo.</li>
  <li>Supports JPEG and PNG image formats.</li>
  <li>Creates a new folder with watermarked images.</li>
</ul>

<h2>Getting Started</h2>

<h3>Prerequisites</h3>

<ul>
  <li>Go installed on your machine.</li>
</ul>

<h3>Installation</h3>

<ol>
  <li>Clone the repository:</li>
</ol>

<pre><code>git clone https://github.com/sepehr-safaeian/Image-Watermark.git
cd Image-Watermark
</code></pre>
<ol>
  <li>Add logo:</li>
</ol>

<pre>
Add your logo to current path
</pre>
<ol start="2">
  <li>Build the executable:</li>
</ol>

<pre><code>go build main.go
</code></pre>

<h3>Usage</h3>

<ol>
  <li>Run the executable:</li>
</ol>

<pre><code>./main
</code></pre>

<ol start="2">
  <li>Enter the path of the folder containing the images when prompted.</li>
  <li>The program will create a new folder named "watermarked" with the watermarked images.</li>
</ol>

<h2>Examples</h2>

<pre><code>Please enter the path of the folder with images: /path/to/images
</code></pre>

<h2>Troubleshooting</h2>

<p>If you encounter any issues, please make sure the specified folder path is correct and the images are in supported formats (JPEG or PNG).</p>

<h2>License</h2>

<p>This project is licensed under the <a href="https://opensource.org/licenses/GPL-3.0">GPL-3.0 License</a> - see the <a href="https://github.com/sepehr-safaeian/Image-Watermark/blob/main/LICENSE">LICENSE</a> file for details.</p>
