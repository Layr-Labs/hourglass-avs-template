# AVS Specifications

This directory contains configuration specifications for your AVS, including runtime values for different environments.

## Runtime Values Configuration

The `runtime/` subdirectory contains environment-specific value overrides for your AVS runtime specifications. These files follow a Helm-style values pattern, allowing you to customize your EigenRuntime specs without modifying templates.

## How it Works

1. **Default Values**: Base configuration is defined in `.hourglass/spec/defaults.yaml`
2. **Environment Overrides**: Place environment-specific overrides in this directory
3. **Deep Merge**: Your values are merged with defaults (your values take precedence)
4. **Template Rendering**: The merged values are used to generate the final runtime specs

## Directory Structure

```
specs/
├── README.md      # This file
└── runtime/       # Environment-specific runtime values
    ├── devnet.yaml    # Development environment overrides
    ├── testnet.yaml   # Testnet environment overrides
    └── mainnet.yaml   # Production environment overrides
```

## Example Configurations

### Development (runtime/devnet.yaml)
```yaml
aggregator:
  env:
    LOG_LEVEL: "debug"
```

### Production (runtime/mainnet.yaml)
```yaml
aggregator:
  env:
    LOG_LEVEL: "warn"
  resources:
    tee_enabled: false

performer:
  env:
    TASK_QUEUE_URL: "https://sqs.us-east-1.amazonaws.com/123456789/prod-tasks"
  resources:
    tee_enabled: true
```

## Available Configuration Options

### Aggregator
```yaml
aggregator:
  env:
    # Any environment variables for the aggregator
    KEY: "value"
  resources:
    tee_enabled: true/false
```

### Executor
```yaml
executor:
  env:
    # Any environment variables for the executor
    KEY: "value"
  resources:
    tee_enabled: true/false
```

### Performer
```yaml
performer:
  env:
    # Any environment variables for the performer
    KEY: "value"
  resources:
    tee_enabled: true/false
```

## Best Practices

1. **Don't Commit Secrets**: Use environment variable references (e.g., `${API_KEY}`) for sensitive data
2. **Environment Parity**: Keep development close to production, but with appropriate debugging settings
3. **Version Control**: Commit these files to track configuration changes over time

## Usage

When you run `devkit avs release publish`, the system automatically:
1. Detects your current context (devnet, testnet, mainnet)
2. Loads the corresponding values file
3. Merges with defaults
4. Generates customized runtime specs
5. Creates and publishes OCI artifacts

No additional flags needed - it just works! 