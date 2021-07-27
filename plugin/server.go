package plugin

import (
	context "context"

	wrapping "github.com/hashicorp/go-kms-wrapping/v2"
	gp "github.com/hashicorp/go-plugin"
	grpc "google.golang.org/grpc"
)

type wrapServer struct {
	UnimplementedWrappingServer
	impl wrapping.Wrapper
}

func (w *wrapper) GRPCServer(broker *gp.GRPCBroker, s *grpc.Server) error {
	RegisterWrappingServer(s, &wrapServer{impl: w.impl})
	return nil
}

func (ws *wrapServer) Type(ctx context.Context, req *TypeRequest) (*TypeResponse, error) {
	return &TypeResponse{Type: uint32(ws.impl.Type(ctx))}, nil
}

func (ws *wrapServer) KeyId(ctx context.Context, req *KeyIdRequest) (*KeyIdResponse, error) {
	return &KeyIdResponse{KeyId: ws.impl.KeyId(ctx)}, nil
}

func (ws *wrapServer) SetConfig(ctx context.Context, req *SetConfigRequest) (*SetConfigResponse, error) {
	opts := req.Options
	if opts == nil {
		opts = new(wrapping.Options)
	}
	wc, err := ws.impl.SetConfig(
		ctx,
		wrapping.WithAad(opts.WithAad),
		wrapping.WithKeyId(opts.WithKeyId),
		wrapping.WithWrapperOptions(opts.WithWrapperOptions),
	)
	if err != nil {
		return nil, err
	}
	return &SetConfigResponse{WrapperConfig: wc}, nil
}

func (ws *wrapServer) Encrypt(ctx context.Context, req *EncryptRequest) (*EncryptResponse, error) {
	opts := req.Options
	if opts == nil {
		opts = new(wrapping.Options)
	}
	ct, err := ws.impl.Encrypt(
		ctx,
		req.Plaintext,
		wrapping.WithAad(opts.WithAad),
		wrapping.WithKeyId(opts.WithKeyId),
		wrapping.WithWrapperOptions(opts.WithWrapperOptions),
	)
	if err != nil {
		return nil, err
	}
	return &EncryptResponse{Ciphertext: ct}, nil
}

func (ws *wrapServer) Decrypt(ctx context.Context, req *DecryptRequest) (*DecryptResponse, error) {
	opts := req.Options
	if opts == nil {
		opts = new(wrapping.Options)
	}
	pt, err := ws.impl.Decrypt(
		ctx,
		req.Ciphertext,
		wrapping.WithAad(opts.WithAad),
		wrapping.WithKeyId(opts.WithKeyId),
		wrapping.WithWrapperOptions(opts.WithWrapperOptions),
	)
	if err != nil {
		return nil, err
	}
	return &DecryptResponse{Plaintext: pt}, nil
}
