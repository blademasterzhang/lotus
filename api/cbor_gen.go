// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package api

import (
	"fmt"
	"io"
	"math"
	"sort"

	abi "github.com/filecoin-project/go-state-types/abi"
	paych "github.com/filecoin-project/specs-actors/actors/builtin/paych"
	market "github.com/filecoin-project/specs-actors/v8/actors/builtin/market"
	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

func (t *PaymentInfo) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{163}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Channel (address.Address) (struct)
	if len("Channel") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Channel\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Channel"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Channel")); err != nil {
		return err
	}

	if err := t.Channel.MarshalCBOR(w); err != nil {
		return err
	}

	// t.WaitSentinel (cid.Cid) (struct)
	if len("WaitSentinel") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"WaitSentinel\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("WaitSentinel"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("WaitSentinel")); err != nil {
		return err
	}

	if err := cbg.WriteCidBuf(scratch, w, t.WaitSentinel); err != nil {
		return xerrors.Errorf("failed to write cid field t.WaitSentinel: %w", err)
	}

	// t.Vouchers ([]*paych.SignedVoucher) (slice)
	if len("Vouchers") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Vouchers\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Vouchers"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Vouchers")); err != nil {
		return err
	}

	if len(t.Vouchers) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Vouchers was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.Vouchers))); err != nil {
		return err
	}
	for _, v := range t.Vouchers {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}
	return nil
}

func (t *PaymentInfo) UnmarshalCBOR(r io.Reader) error {
	*t = PaymentInfo{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("PaymentInfo: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringBuf(br, scratch)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Channel (address.Address) (struct)
		case "Channel":

			{

				if err := t.Channel.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("unmarshaling t.Channel: %w", err)
				}

			}
			// t.WaitSentinel (cid.Cid) (struct)
		case "WaitSentinel":

			{

				c, err := cbg.ReadCid(br)
				if err != nil {
					return xerrors.Errorf("failed to read cid field t.WaitSentinel: %w", err)
				}

				t.WaitSentinel = c

			}
			// t.Vouchers ([]*paych.SignedVoucher) (slice)
		case "Vouchers":

			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return err
			}

			if extra > cbg.MaxLength {
				return fmt.Errorf("t.Vouchers: array too large (%d)", extra)
			}

			if maj != cbg.MajArray {
				return fmt.Errorf("expected cbor array")
			}

			if extra > 0 {
				t.Vouchers = make([]*paych.SignedVoucher, extra)
			}

			for i := 0; i < int(extra); i++ {

				var v paych.SignedVoucher
				if err := v.UnmarshalCBOR(br); err != nil {
					return err
				}

				t.Vouchers[i] = &v
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *SealedRef) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{163}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.SectorID (abi.SectorNumber) (uint64)
	if len("SectorID") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"SectorID\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("SectorID"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("SectorID")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.SectorID)); err != nil {
		return err
	}

	// t.Offset (abi.PaddedPieceSize) (uint64)
	if len("Offset") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Offset\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Offset"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Offset")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Offset)); err != nil {
		return err
	}

	// t.Size (abi.UnpaddedPieceSize) (uint64)
	if len("Size") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Size\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Size"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Size")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Size)); err != nil {
		return err
	}

	return nil
}

func (t *SealedRef) UnmarshalCBOR(r io.Reader) error {
	*t = SealedRef{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("SealedRef: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringBuf(br, scratch)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.SectorID (abi.SectorNumber) (uint64)
		case "SectorID":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.SectorID = abi.SectorNumber(extra)

			}
			// t.Offset (abi.PaddedPieceSize) (uint64)
		case "Offset":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.Offset = abi.PaddedPieceSize(extra)

			}
			// t.Size (abi.UnpaddedPieceSize) (uint64)
		case "Size":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.Size = abi.UnpaddedPieceSize(extra)

			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *SealedRefs) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{161}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Refs ([]api.SealedRef) (slice)
	if len("Refs") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Refs\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Refs"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Refs")); err != nil {
		return err
	}

	if len(t.Refs) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Refs was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.Refs))); err != nil {
		return err
	}
	for _, v := range t.Refs {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}
	return nil
}

func (t *SealedRefs) UnmarshalCBOR(r io.Reader) error {
	*t = SealedRefs{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("SealedRefs: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringBuf(br, scratch)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Refs ([]api.SealedRef) (slice)
		case "Refs":

			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return err
			}

			if extra > cbg.MaxLength {
				return fmt.Errorf("t.Refs: array too large (%d)", extra)
			}

			if maj != cbg.MajArray {
				return fmt.Errorf("expected cbor array")
			}

			if extra > 0 {
				t.Refs = make([]SealedRef, extra)
			}

			for i := 0; i < int(extra); i++ {

				var v SealedRef
				if err := v.UnmarshalCBOR(br); err != nil {
					return err
				}

				t.Refs[i] = v
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *SealTicket) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{162}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Value (abi.SealRandomness) (slice)
	if len("Value") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Value\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Value"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Value")); err != nil {
		return err
	}

	if len(t.Value) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Value was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Value))); err != nil {
		return err
	}

	if _, err := w.Write(t.Value[:]); err != nil {
		return err
	}

	// t.Epoch (abi.ChainEpoch) (int64)
	if len("Epoch") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Epoch\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Epoch"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Epoch")); err != nil {
		return err
	}

	if t.Epoch >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Epoch)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.Epoch-1)); err != nil {
			return err
		}
	}
	return nil
}

func (t *SealTicket) UnmarshalCBOR(r io.Reader) error {
	*t = SealTicket{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("SealTicket: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringBuf(br, scratch)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Value (abi.SealRandomness) (slice)
		case "Value":

			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return err
			}

			if extra > cbg.ByteArrayMaxLen {
				return fmt.Errorf("t.Value: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}

			if extra > 0 {
				t.Value = make([]uint8, extra)
			}

			if _, err := io.ReadFull(br, t.Value[:]); err != nil {
				return err
			}
			// t.Epoch (abi.ChainEpoch) (int64)
		case "Epoch":
			{
				maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
				var extraI int64
				if err != nil {
					return err
				}
				switch maj {
				case cbg.MajUnsignedInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 positive overflow")
					}
				case cbg.MajNegativeInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 negative oveflow")
					}
					extraI = -1 - extraI
				default:
					return fmt.Errorf("wrong type for int64 field: %d", maj)
				}

				t.Epoch = abi.ChainEpoch(extraI)
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *SealSeed) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{162}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Value (abi.InteractiveSealRandomness) (slice)
	if len("Value") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Value\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Value"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Value")); err != nil {
		return err
	}

	if len(t.Value) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Value was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Value))); err != nil {
		return err
	}

	if _, err := w.Write(t.Value[:]); err != nil {
		return err
	}

	// t.Epoch (abi.ChainEpoch) (int64)
	if len("Epoch") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Epoch\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Epoch"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Epoch")); err != nil {
		return err
	}

	if t.Epoch >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Epoch)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.Epoch-1)); err != nil {
			return err
		}
	}
	return nil
}

func (t *SealSeed) UnmarshalCBOR(r io.Reader) error {
	*t = SealSeed{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("SealSeed: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringBuf(br, scratch)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Value (abi.InteractiveSealRandomness) (slice)
		case "Value":

			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return err
			}

			if extra > cbg.ByteArrayMaxLen {
				return fmt.Errorf("t.Value: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}

			if extra > 0 {
				t.Value = make([]uint8, extra)
			}

			if _, err := io.ReadFull(br, t.Value[:]); err != nil {
				return err
			}
			// t.Epoch (abi.ChainEpoch) (int64)
		case "Epoch":
			{
				maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
				var extraI int64
				if err != nil {
					return err
				}
				switch maj {
				case cbg.MajUnsignedInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 positive overflow")
					}
				case cbg.MajNegativeInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 negative oveflow")
					}
					extraI = -1 - extraI
				default:
					return fmt.Errorf("wrong type for int64 field: %d", maj)
				}

				t.Epoch = abi.ChainEpoch(extraI)
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *PieceDealInfo) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{165}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.PublishCid (cid.Cid) (struct)
	if len("PublishCid") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"PublishCid\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("PublishCid"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("PublishCid")); err != nil {
		return err
	}

	if t.PublishCid == nil {
		if _, err := w.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCidBuf(scratch, w, *t.PublishCid); err != nil {
			return xerrors.Errorf("failed to write cid field t.PublishCid: %w", err)
		}
	}

	// t.DealID (abi.DealID) (uint64)
	if len("DealID") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"DealID\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("DealID"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("DealID")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.DealID)); err != nil {
		return err
	}

	// t.DealProposal (market.DealProposal) (struct)
	if len("DealProposal") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"DealProposal\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("DealProposal"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("DealProposal")); err != nil {
		return err
	}

	if err := t.DealProposal.MarshalCBOR(w); err != nil {
		return err
	}

	// t.DealSchedule (api.DealSchedule) (struct)
	if len("DealSchedule") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"DealSchedule\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("DealSchedule"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("DealSchedule")); err != nil {
		return err
	}

	if err := t.DealSchedule.MarshalCBOR(w); err != nil {
		return err
	}

	// t.KeepUnsealed (bool) (bool)
	if len("KeepUnsealed") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"KeepUnsealed\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("KeepUnsealed"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("KeepUnsealed")); err != nil {
		return err
	}

	if err := cbg.WriteBool(w, t.KeepUnsealed); err != nil {
		return err
	}
	return nil
}

func (t *PieceDealInfo) UnmarshalCBOR(r io.Reader) error {
	*t = PieceDealInfo{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("PieceDealInfo: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringBuf(br, scratch)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.PublishCid (cid.Cid) (struct)
		case "PublishCid":

			{

				b, err := br.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := br.UnreadByte(); err != nil {
						return err
					}

					c, err := cbg.ReadCid(br)
					if err != nil {
						return xerrors.Errorf("failed to read cid field t.PublishCid: %w", err)
					}

					t.PublishCid = &c
				}

			}
			// t.DealID (abi.DealID) (uint64)
		case "DealID":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.DealID = abi.DealID(extra)

			}
			// t.DealProposal (market.DealProposal) (struct)
		case "DealProposal":

			{

				b, err := br.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := br.UnreadByte(); err != nil {
						return err
					}
					t.DealProposal = new(market.DealProposal)
					if err := t.DealProposal.UnmarshalCBOR(br); err != nil {
						return xerrors.Errorf("unmarshaling t.DealProposal pointer: %w", err)
					}
				}

			}
			// t.DealSchedule (api.DealSchedule) (struct)
		case "DealSchedule":

			{

				if err := t.DealSchedule.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("unmarshaling t.DealSchedule: %w", err)
				}

			}
			// t.KeepUnsealed (bool) (bool)
		case "KeepUnsealed":

			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return err
			}
			if maj != cbg.MajOther {
				return fmt.Errorf("booleans must be major type 7")
			}
			switch extra {
			case 20:
				t.KeepUnsealed = false
			case 21:
				t.KeepUnsealed = true
			default:
				return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *DealSchedule) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{162}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.StartEpoch (abi.ChainEpoch) (int64)
	if len("StartEpoch") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"StartEpoch\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("StartEpoch"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("StartEpoch")); err != nil {
		return err
	}

	if t.StartEpoch >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.StartEpoch)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.StartEpoch-1)); err != nil {
			return err
		}
	}

	// t.EndEpoch (abi.ChainEpoch) (int64)
	if len("EndEpoch") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"EndEpoch\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("EndEpoch"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("EndEpoch")); err != nil {
		return err
	}

	if t.EndEpoch >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.EndEpoch)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.EndEpoch-1)); err != nil {
			return err
		}
	}
	return nil
}

func (t *DealSchedule) UnmarshalCBOR(r io.Reader) error {
	*t = DealSchedule{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("DealSchedule: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringBuf(br, scratch)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.StartEpoch (abi.ChainEpoch) (int64)
		case "StartEpoch":
			{
				maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
				var extraI int64
				if err != nil {
					return err
				}
				switch maj {
				case cbg.MajUnsignedInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 positive overflow")
					}
				case cbg.MajNegativeInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 negative oveflow")
					}
					extraI = -1 - extraI
				default:
					return fmt.Errorf("wrong type for int64 field: %d", maj)
				}

				t.StartEpoch = abi.ChainEpoch(extraI)
			}
			// t.EndEpoch (abi.ChainEpoch) (int64)
		case "EndEpoch":
			{
				maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
				var extraI int64
				if err != nil {
					return err
				}
				switch maj {
				case cbg.MajUnsignedInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 positive overflow")
					}
				case cbg.MajNegativeInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 negative oveflow")
					}
					extraI = -1 - extraI
				default:
					return fmt.Errorf("wrong type for int64 field: %d", maj)
				}

				t.EndEpoch = abi.ChainEpoch(extraI)
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
