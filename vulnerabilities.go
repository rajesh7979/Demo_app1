package denySecurityVulnerabilities

deny_vulnerabilities {
    input.action == "deploy"  # Modify action as needed for your pipeline trigger

    # Extract Helm chart vulnerabilities from input
    helmChart := input.request.body.input.helmChart
    vulnerabilities := helmChart.spec.vulnerabilities

    # Check if vulnerabilities exist and deny the deployment if found
    count(vulnerabilities) > 0
}
