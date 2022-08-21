FROM temp:latest as base_img

# ------------------------------------------------------------------------------
# Production image
# ------------------------------------------------------------------------------
FROM alpine:latest
WORKDIR /work
COPY --from=base_img /work/oidc /work/oidc
EXPOSE 9998
ENTRYPOINT /work/oidc