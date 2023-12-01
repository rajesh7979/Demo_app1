package denyMismatchedImage

deny_mismatched_image {
    input.action == "deploy"  # Adjust action as per your pipeline trigger

    # Extract image from the pipeline input
    image := input.request.body.input.image

    # Extract image from Helm chart
    helmChart := input.request.body.input.helmChart
    chartImage := helmChart.spec.image

    # Deny if the images don't match
    not startswith(image, chartImage)
}
