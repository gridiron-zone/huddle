---
id: full-sync
title: Full Sync
sidebar_label: Full Sync
slug: full-sync
---
# Mainnet Full sync
## Software downgrade 

:::note   
You will need to build the first version of the Huddle mainnet in order to be able to sync the chain from scratch.
:::

```bash
# Make sure we are inside the home directory
cd $HOME

# Clone the Huddle software
git clone https://github.com/gridiron-zone/huddle.git && cd huddle

# Checkout the correct tag
git checkout tags/v1.0.3

# Build the software
# If you want to use the default database backend run
make install
```
