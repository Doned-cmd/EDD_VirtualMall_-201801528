import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AnioPedidoComponent } from './anio-pedido.component';

describe('AnioPedidoComponent', () => {
  let component: AnioPedidoComponent;
  let fixture: ComponentFixture<AnioPedidoComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AnioPedidoComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AnioPedidoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
