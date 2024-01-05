# Feature

## Compress use case

### 1. Specify target image files

* Specify the target image files **by directory**, and the destination directory.
* Additionally, specify the dimension (width \* height) to scale down.

### 2. Compress

* Scale down: scale down the input image to the specified dimension.
* Convert image file format: to `.webp`.
* Reduce the quality: determine the quality with reference to SSIM.

### 3. Check the results

Display the following

* files that have changed
* SSIM
* file size (data compression ratio)
